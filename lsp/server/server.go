package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	laravel_ls "laravel-ls"
	"laravel-ls/cache"
	"laravel-ls/lsp/protocol"
	"laravel-ls/provider"

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

	providerManager *provider.Manager
}

func NewServer(providerManager *provider.Manager) *Server {
	return &Server{
		cache:           cache.NewFileCache(),
		providerManager: providerManager,
	}
}

func validateURI(uri string) (string, error) {
	var err error = nil
	path, ok := strings.CutPrefix(uri, "file://")
	if !ok {
		err = ErrNonLocalPath
	}
	return path, err
}

func (s *Server) HandleTextDocumentCompletion(params protocol.CompletionParams) (protocol.CompletionResult, error) {
	log.WithField("method", protocol.MethodTextDocumentCompletion).
		WithField("filename", params.TextDocument.URI).
		Info("completion")

	response := protocol.CompletionResult{
		CompletionItems: []protocol.CompletionItem{},
	}

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return response, err
	}

	file := s.cache.Get(filename)

	context := provider.CompletionContext{
		BaseContext: provider.BaseContext{
			Logger: log.WithField("module", "Definition"),
			File:   file,
		},
		Position: toTSPoint(params.Position),
		Publish: func(item protocol.CompletionItem) {
			response.CompletionItems = append(response.CompletionItems, item)
		},
	}

	s.providerManager.Completion(context)

	return response, err
}

func (s *Server) HandleTextDocumentHover(params protocol.HoverParams) (protocol.HoverResult, error) {
	log.WithField("method", protocol.MethodTextDocumentHover).
		WithField("filename", params.TextDocument.URI).
		Info("Hover")

	response := protocol.HoverResult{}

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return response, err
	}

	file := s.cache.Get(filename)

	context := provider.HoverContext{
		BaseContext: provider.BaseContext{
			Logger:    log.WithField("module", "Definition"),
			File:      file,
			FileCache: s.cache,
		},
		Position: toTSPoint(params.Position),
	}

	content := s.providerManager.Hover(context)
	if len(content) > 0 {
		response.Contents.MarkupContent = &protocol.MarkupContent{
			Kind:  "markup",
			Value: content,
		}
	}
	return response, nil
}

func (s *Server) HandleTextDocumentDiagnostic(params protocol.DocumentDiagnosticParams) (protocol.DocumentDiagnosticReport, error) {
	log.WithField("method", protocol.MethodTextDocumentDiagnostic).
		WithField("filename", params.TextDocument.URI).
		Info("Diagnostic")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return &protocol.FullDocumentDiagnosticReport{}, err
	}

	file := s.cache.Get(filename)
	if file == nil {
		return &protocol.FullDocumentDiagnosticReport{}, ErrFileNotOpened
	}

	items := []protocol.Diagnostic{}
	for _, diagnostic := range s.providerManager.Diagnostics(file) {

		start := diagnostic.Range.StartPoint
		end := diagnostic.Range.EndPoint

		items = append(items, protocol.Diagnostic{
			Range: protocol.Range{
				Start: FromTSPoint(start),
				End:   FromTSPoint(end),
			},
			Severity: protocol.DiagnosticSeverity(diagnostic.Severity),
			Source:   laravel_ls.Name,
			Message:  diagnostic.Message,
		})
	}

	return &protocol.FullDocumentDiagnosticReport{
		Kind:  "full",
		Items: items,
	}, nil
}

func (s *Server) HandleTextDocumentDefinition(params protocol.DefinitionParams) (response protocol.DefinitionResponse, err error) {
	log.WithField("method", protocol.MethodTextDocumentDefinition).
		WithField("filename", params.TextDocument.URI).
		Info("Definition")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return response, err
	}

	file := s.cache.Get(filename)
	if file == nil {
		return response, ErrFileNotOpened
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
			response.Locations = append(response.Locations, location)
		},
	}

	s.providerManager.ResolveDefinition(context)

	return
}

func (s Server) HandleTextDocumentDidOpen(params protocol.DidOpenTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidOpen).
		WithField("lang", params.TextDocument.LanguageID).
		WithField("filename", params.TextDocument.URI).
		Info("Document opened")

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
		Info("Document changed")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return err
	}

	file := s.cache.Get(filename)
	if file == nil {
		return ErrFileNotOpened
	}

	var errs error = nil

	for _, change := range params.ContentChanges {

		start := toTSPoint(change.Range.Start)
		end := toTSPoint(change.Range.End)

		log.Info("Change", start, end, change.Text)

		err := file.Update(start, end, []byte(change.Text))
		if err != nil {
			errs = errors.Join(err)
		}
	}

	return errs
}

func (s Server) HandleTextDocumentDidSave(params protocol.DidSaveTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidSave).
		WithField("filename", params.TextDocument.URI).
		Info("Document saved")
	return nil
}

func (s Server) HandleTextDocumentDidClose(params protocol.DidCloseTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidClose).
		WithField("filename", params.TextDocument.URI).
		Info("Document closed")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return err
	}

	return s.cache.Close(filename)
}

func (s *Server) HandleInitialize(params protocol.InitializeParams) (protocol.InitializeResult, error) {
	rootPath, found := strings.CutPrefix(params.RootURI, "file://")
	if !found {
		return protocol.InitializeResult{}, fmt.Errorf("server only support local filesystem root paths")
	}

	log.WithField("method", protocol.MethodInitialize).
		WithField("rootPath", rootPath).
		Info("Initialized")

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
			DiagnosticProvider: true,
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    laravel_ls.Name,
			Version: laravel_ls.Version,
		},
	}, nil
}

// Handle incoming LSP messages
func (s *Server) dispatch(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {
	switch req.Method {
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
			Info("Initialized")
		return nil, nil
	default:
		// Respond with a method not found error
		return nil, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: fmt.Sprintf("Method %s not found", req.Method),
		}
	}
}

func (s Server) Run(ctx context.Context, conn io.ReadWriteCloser) error {
	stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{})
	rpc := jsonrpc2.NewConn(ctx, stream, jsonrpc2.HandlerWithError(s.dispatch))

	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	case <-rpc.DisconnectNotify():
		return nil
	}
}
