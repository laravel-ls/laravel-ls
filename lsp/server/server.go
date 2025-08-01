package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/laravel-ls/laravel-ls/cache"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/program"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/uri"

	log "github.com/sirupsen/logrus"
	"github.com/sourcegraph/jsonrpc2"
)

var (
	ErrNonLocalPath             = errors.New("server only support local filesystem paths")
	ErrFileNotOpened            = errors.New("file not opened")
	ErrFailedToGetPointAtCursor = errors.New("failed to get node at cursor")
)

type Server struct {
	// Map of open files for this session
	cache *cache.FileCache

	// flag if shutdown request has been received.
	// if a connection is closed without this request, it is an error.
	shutdownReceived bool

	providerManager *provider.Manager
}

func NewServer(providerManager *provider.Manager) *Server {
	return &Server{
		cache:           cache.NewFileCache(),
		providerManager: providerManager,
	}
}

func validateURI(input string) (string, error) {
	u, err := uri.Parse(input)
	if err != nil {
		return "", err
	}
	if !u.HasFilename() {
		return "", ErrNonLocalPath
	}

	return u.Filename(), nil
}

func (s *Server) HandleTextDocumentCodeAction(params protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	log.WithField("method", protocol.MethodTextDocumentCodeAction).
		WithField("filename", params.TextDocument.URI).
		Info("code action")

	response := []protocol.CodeAction{}

	file, err := s.getFile(params.TextDocument)
	if err != nil {
		return response, err
	}

	s.providerManager.CodeAction(provider.CodeActionContext{
		BaseContext: provider.BaseContext{
			Logger:    log.WithField("module", "CodeAction"),
			File:      file,
			FileCache: s.cache,
		},
		Range: toTSRange(params.Range),
		Publish: func(codeAction protocol.CodeAction) {
			response = append(response, codeAction)
		},
	})

	return response, nil
}

func (s *Server) HandleTextDocumentCompletion(params protocol.CompletionParams) (protocol.CompletionResponse, error) {
	log.WithField("method", protocol.MethodTextDocumentCompletion).
		WithField("filename", params.TextDocument.URI).
		Debug("completion")

	response := protocol.CompletionResponse{
		Items: []protocol.CompletionItem{},
	}

	file, err := s.getFile(params.TextDocument)
	if err != nil {
		return response, ErrFileNotOpened
	}

	context := provider.CompletionContext{
		BaseContext: provider.BaseContext{
			Logger:    log.WithField("module", "Definition"),
			File:      file,
			FileCache: s.cache,
		},
		Position: toTSPoint(params.Position),
		Publish: func(item protocol.CompletionItem) {
			response.Items = append(response.Items, item)
		},
	}

	s.providerManager.Completion(context)

	return response, err
}

func (s *Server) HandleTextDocumentHover(params protocol.HoverParams) (protocol.HoverResult, error) {
	log.WithField("method", protocol.MethodTextDocumentHover).
		WithField("filename", params.TextDocument.URI).
		Debug("Hover")

	response := protocol.HoverResult{}

	file, err := s.getFile(params.TextDocument)
	if err != nil {
		return response, err
	}

	content := ""

	s.providerManager.Hover(provider.HoverContext{
		BaseContext: provider.BaseContext{
			Logger:    log.WithField("module", "Definition"),
			File:      file,
			FileCache: s.cache,
		},
		Position: toTSPoint(params.Position),
		Publish: func(result provider.Hover) {
			content += result.Content
		},
	})

	if len(content) > 0 {
		response.Hover = &protocol.Hover{
			Contents: protocol.MarkupContentOrMarkedString{
				Markup: &protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: content,
				},
			},
		}
	}
	return response, nil
}

func (s *Server) HandleTextDocumentDiagnostic(params protocol.DocumentDiagnosticParams) (protocol.DocumentDiagnosticReport, error) {
	log.WithField("method", protocol.MethodTextDocumentDiagnostic).
		WithField("filename", params.TextDocument.URI).
		Debug("Diagnostic")

	file, err := s.getFile(params.TextDocument)
	if file == nil {
		return protocol.DocumentDiagnosticReport{}, err
	}

	items := []protocol.Diagnostic{}

	s.providerManager.Diagnostics(provider.DiagnosticContext{
		BaseContext: provider.BaseContext{
			Logger:    log.WithField("module", "diagnostic"),
			File:      file,
			FileCache: s.cache,
		},
		Publish: func(diagnostic provider.Diagnostic) {
			start := diagnostic.Range.StartPoint
			end := diagnostic.Range.EndPoint

			items = append(items, protocol.Diagnostic{
				Range: protocol.Range{
					Start: FromTSPoint(start),
					End:   FromTSPoint(end),
				},
				Severity: diagnostic.Severity,
				Source:   program.Name,
				Message:  diagnostic.Message,
			})
		},
	})

	return protocol.DocumentDiagnosticReport{
		Full: &protocol.FullDocumentDiagnosticReport{
			Kind:  "full",
			Items: items,
		},
	}, nil
}

func (s *Server) HandleTextDocumentDefinition(params protocol.DefinitionParams) (response protocol.DefinitionResponse, err error) {
	log.WithField("method", protocol.MethodTextDocumentDefinition).
		WithField("filename", params.TextDocument.URI).
		Debug("Definition")

	file, err := s.getFile(params.TextDocument)
	if err != nil {
		return response, err
	}

	logger := log.WithField("module", "Definition")

	context := provider.DefinitionContext{
		BaseContext: provider.BaseContext{
			Logger:    logger,
			FileCache: s.cache,
			File:      file,
		},
		Position: toTSPoint(params.Position),
		Publish: func(location protocol.Location) {
			location.URI = "file://" + location.URI
			response.LocationList = append(response.LocationList, location)
		},
	}

	s.providerManager.ResolveDefinition(context)

	return
}

func (s Server) HandleTextDocumentDidOpen(params protocol.DidOpenTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidOpen).
		WithField("lang", params.TextDocument.LanguageID).
		WithField("filename", params.TextDocument.URI).
		Debug("Document opened")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return err
	}

	_, err = s.cache.Open(filename)
	return err
}

func (s Server) HandleTextDocumentDidChange(params protocol.DidChangeTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidChange).
		WithField("filename", params.TextDocument.URI).
		Debug("Document changed")

	file, err := s.getFile(params.TextDocument.TextDocumentIdentifier)
	if err != nil {
		return ErrFileNotOpened
	}

	var errs error = nil

	for _, change := range params.ContentChanges {

		start := toTSPoint(change.Range.Start)
		end := toTSPoint(change.Range.End)

		log.Debug("Change", start, end, change.Text)

		err := file.Update(start, end, []byte(change.Text))
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

func (s Server) HandleTextDocumentDidSave(params protocol.DidSaveTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidSave).
		WithField("filename", params.TextDocument.URI).
		Debug("Document saved")
	return nil
}

func (s Server) HandleTextDocumentDidClose(params protocol.DidCloseTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidClose).
		WithField("filename", params.TextDocument.URI).
		Debug("Document closed")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return err
	}

	return s.cache.Close(filename)
}

func (s *Server) HandleInitialize(params protocol.InitializeParams) (protocol.InitializeResult, error) {
	rootPath, err := validateURI(string(params.RootURI))
	if err == ErrNonLocalPath {
		return protocol.InitializeResult{}, fmt.Errorf("server only support local filesystem root paths")
	} else if err != nil {
		return protocol.InitializeResult{}, err
	}

	log.WithField("method", protocol.MethodInitialize).
		WithField("rootPath", rootPath).
		Debug("Initialize")

	s.providerManager.Init(provider.InitContext{
		Logger:    log.WithField("module", "Initialize"),
		RootPath:  rootPath,
		FileCache: s.cache,
	})

	// Respond with capabilities
	return protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncKindIncremental,
			HoverProvider:    true,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: []string{"'", "\""},
			},
			DefinitionProvider: true,
			DiagnosticProvider: protocol.DiagnosticOptions{
				InterFileDependencies: true,
				WorkspaceDiagnostics:  false,
			},
			CodeActionProvider: true,
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    program.Name,
			Version: program.Version(),
		},
	}, nil
}

// Handle incoming LSP messages
func (s *Server) dispatch(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {
	switch req.Method {
	case protocol.MethodTextDocumentCodeAction:
		var params protocol.CodeActionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.HandleTextDocumentCodeAction(params)
	case protocol.MethodTextDocumentCompletion:
		var params protocol.CompletionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.HandleTextDocumentCompletion(params)
	case protocol.MethodTextDocumentHover:
		var params protocol.HoverParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.HandleTextDocumentHover(params)
	case protocol.MethodTextDocumentDiagnostic:
		var params protocol.DocumentDiagnosticParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.HandleTextDocumentDiagnostic(params)
	case protocol.MethodTextDocumentDefinition:
		var params protocol.DefinitionParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.HandleTextDocumentDefinition(params)
	case protocol.MethodTextDocumentDidOpen:
		var params protocol.DidOpenTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return nil, s.HandleTextDocumentDidOpen(params)
	case protocol.MethodTextDocumentDidChange:
		var params protocol.DidChangeTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return nil, s.HandleTextDocumentDidChange(params)
	case protocol.MethodTextDocumentDidSave:
		var params protocol.DidSaveTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return nil, s.HandleTextDocumentDidSave(params)
	case protocol.MethodTextDocumentDidClose:
		var params protocol.DidCloseTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return nil, s.HandleTextDocumentDidClose(params)
	case protocol.MethodInitialize:
		var params protocol.InitializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return s.HandleInitialize(params)
	case protocol.MethodInitialized:
		log.WithField("method", protocol.MethodInitialized).
			Debug("Initialized")
		return nil, nil
	case "$/cancelRequest":
		// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#cancelRequest
		// TODO: Maybe implement a way to cancel requests?
		return nil, nil
	case "shutdown":
		// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#shutdown
		// TODO: Implement shutdown logic if needed - ie. clear temp files, close connections, etc.
		log.Info("Received shutdown request")
		s.shutdownReceived = true
		return nil, nil
	case "exit":
		// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#exit
		log.Info("Received exit notification")
		conn.Close()
		return nil, nil
	default:
		log.WithField("method", req.Method).Warn("LSP method not found")
		return nil, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: fmt.Sprintf("Method %s not found", req.Method),
		}
	}
}

type jsonRPCLogger struct{}

func (jsonRPCLogger) Printf(format string, v ...any) {
	log.Tracef(format, v...)
}

func (s Server) Run(ctx context.Context, conn io.ReadWriteCloser) error {
	stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{})

	opts := []jsonrpc2.ConnOpt{}
	if log.GetLevel() >= log.TraceLevel {
		opts = append(opts, jsonrpc2.LogMessages(jsonRPCLogger{}))
	}
	log.Info("Started laravel-ls server")
	rpc := jsonrpc2.NewConn(ctx, stream, jsonrpc2.HandlerWithError(s.dispatch), opts...)

	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	case <-rpc.DisconnectNotify():
		if !s.shutdownReceived {
			return fmt.Errorf("disconnected without an shutdown request")
		}
		return nil
	}
}

func (s *Server) getFile(identifier protocol.TextDocumentIdentifier) (*parser.File, error) {
	filename, err := validateURI(identifier.URI)
	if err != nil {
		return nil, err
	}

	file := s.cache.Get(filename)
	if file == nil {
		return nil, ErrFileNotOpened
	}
	return file, nil
}
