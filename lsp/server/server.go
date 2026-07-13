package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/laravel-ls/laravel-ls/cache"
	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/program"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/protocol"
	"github.com/laravel-ls/uri"

	jsonrpc "github.com/gumeniukcom/golang-jsonrpc2/v2"
	"github.com/gumeniukcom/golang-jsonrpc2/v2/jsonrpcstdio"
	log "github.com/sirupsen/logrus"
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

	// config is parsed from initializationOptions during the initialize handshake.
	config LSPConfig
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

	return response, err
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

func (s *Server) HandleTextDocumentDidSave(params protocol.DidSaveTextDocumentParams) error {
	log.WithField("method", protocol.MethodTextDocumentDidSave).
		WithField("filename", params.TextDocument.URI).
		Debug("Document saved")

	filename, err := validateURI(params.TextDocument.URI)
	if err != nil {
		return err
	}

	// Fire-and-forget: the returned channel closes when providers finish
	// re-warming, but the server does not wait for it. Callers that need to
	// synchronise on cache readiness should await the channel themselves.
	_ = s.providerManager.FileSaved(filename)

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

	if params.InitializationOptions != nil {
		raw, err := json.Marshal(params.InitializationOptions)
		if err != nil {
			log.WithError(err).Warn("failed to marshal initializationOptions")
		} else if err := json.Unmarshal(raw, &s.config); err != nil {
			log.WithError(err).Warn("failed to parse initializationOptions")
		}
	}

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
			InlayHintProvider:  s.config.InlayHints.Routes.IsEnabled(),
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    program.Name,
			Version: program.Version(),
		},
	}, nil
}

func (s *Server) HandleTextDocumentInlayHint(params protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	log.WithField("method", protocol.MethodTextDocumentInlayHint).
		WithField("filename", params.TextDocument.URI).
		Debug("InlayHint")

	response := []protocol.InlayHint{}

	if !s.config.InlayHints.Routes.IsEnabled() {
		return response, nil
	}

	file, err := s.getFile(params.TextDocument)
	if err != nil {
		return response, err
	}

	s.providerManager.InlayHints(provider.InlayHintContext{
		BaseContext: provider.BaseContext{
			Logger:    log.WithField("module", "InlayHint"),
			File:      file,
			FileCache: s.cache,
		},
		Range: toTSRange(params.Range),
		Publish: func(hint provider.InlayHint) {
			paddingLeft := true
			response = append(response, protocol.InlayHint{
				Position:    FromTSPoint(hint.Position),
				PaddingLeft: &paddingLeft,
				Label: protocol.InlayHintLabel{
					String: &hint.Label,
				},
			})
		},
	})

	return response, nil
}

// registerMethods wires every LSP method onto the dispatcher. Typed
// registration replaces the previous hand-written dispatch switch: params
// unmarshaling, routing, and method-not-found responses are handled by the
// library. exitFn terminates the session (LSP "exit").
func (s *Server) registerMethods(rpc *jsonrpc.JSONRPC, exitFn func()) error {
	return errors.Join(
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentCodeAction,
			func(_ context.Context, p protocol.CodeActionParams) ([]protocol.CodeAction, error) {
				return s.HandleTextDocumentCodeAction(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentCompletion,
			func(_ context.Context, p protocol.CompletionParams) (protocol.CompletionResponse, error) {
				return s.HandleTextDocumentCompletion(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentHover,
			func(_ context.Context, p protocol.HoverParams) (protocol.HoverResult, error) {
				return s.HandleTextDocumentHover(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentDiagnostic,
			func(_ context.Context, p protocol.DocumentDiagnosticParams) (protocol.DocumentDiagnosticReport, error) {
				return s.HandleTextDocumentDiagnostic(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentDefinition,
			func(_ context.Context, p protocol.DefinitionParams) (protocol.DefinitionResponse, error) {
				return s.HandleTextDocumentDefinition(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentInlayHint,
			func(_ context.Context, p protocol.InlayHintParams) ([]protocol.InlayHint, error) {
				return s.HandleTextDocumentInlayHint(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodInitialize,
			func(_ context.Context, p protocol.InitializeParams) (protocol.InitializeResult, error) {
				return s.HandleInitialize(p)
			}),

		// Notifications: executed, never answered (per the JSON-RPC spec).
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentDidOpen,
			func(_ context.Context, p protocol.DidOpenTextDocumentParams) (struct{}, error) {
				return struct{}{}, s.HandleTextDocumentDidOpen(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentDidChange,
			func(_ context.Context, p protocol.DidChangeTextDocumentParams) (struct{}, error) {
				return struct{}{}, s.HandleTextDocumentDidChange(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentDidSave,
			func(_ context.Context, p protocol.DidSaveTextDocumentParams) (struct{}, error) {
				return struct{}{}, s.HandleTextDocumentDidSave(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodTextDocumentDidClose,
			func(_ context.Context, p protocol.DidCloseTextDocumentParams) (struct{}, error) {
				return struct{}{}, s.HandleTextDocumentDidClose(p)
			}),
		jsonrpc.RegisterTyped(rpc, protocol.MethodInitialized,
			func(_ context.Context, _ json.RawMessage) (struct{}, error) {
				log.WithField("method", protocol.MethodInitialized).Debug("Initialized")
				return struct{}{}, nil
			}),

		// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#cancelRequest
		// TODO: Maybe implement a way to cancel requests?
		jsonrpc.RegisterTyped(rpc, "$/cancelRequest",
			func(_ context.Context, _ json.RawMessage) (struct{}, error) {
				return struct{}{}, nil
			}),

		// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#shutdown
		// TODO: Implement shutdown logic if needed - ie. clear temp files, close connections, etc.
		jsonrpc.RegisterTyped(rpc, "shutdown",
			func(_ context.Context, _ json.RawMessage) (*struct{}, error) {
				log.Info("Received shutdown request")
				s.shutdownReceived = true
				return nil, nil // LSP expects a null result
			}),

		// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#exit
		jsonrpc.RegisterTyped(rpc, "exit",
			func(_ context.Context, _ json.RawMessage) (struct{}, error) {
				log.Info("Received exit notification")
				exitFn()
				return struct{}{}, nil
			}),
	)
}

func (s *Server) Run(ctx context.Context, conn io.ReadWriteCloser) error {
	rpc := jsonrpc.New()
	// Error texts never reach the client (the library answers with generic
	// codes and keeps detail server-side); route that detail through logrus
	// instead of the library's slog logger so the log stream stays uniform.
	rpc.SetLogger(nil)
	rpc.Use(func(method string, next jsonrpc.RPCMethod) jsonrpc.RPCMethod {
		return func(ctx context.Context, data json.RawMessage) (json.RawMessage, int, error) {
			res, code, err := next(ctx, data)
			if err != nil {
				log.WithField("method", method).WithError(err).Warn("handler error")
			}
			return res, code, err
		}
	})
	// Methods the dispatcher rejects before any handler runs (unknown
	// method, invalid params) never reach middleware; keep the old Warn for
	// unknown methods via the observability hook.
	rpc.SetObserver(func(_ context.Context, info jsonrpc.CallInfo) {
		if info.Code == jsonrpc.MethodNotFoundErrorCode {
			log.WithField("method", info.Method).Warn("LSP method not found")
		}
	})
	// The previous implementation had no per-request timeout; the library
	// defaults to 30s, so raise it to a generous safety bound.
	rpc.SetDefaultTimeOut(15 * time.Minute)

	if log.GetLevel() >= log.TraceLevel {
		rpc.Use(func(method string, next jsonrpc.RPCMethod) jsonrpc.RPCMethod {
			return func(ctx context.Context, data json.RawMessage) (json.RawMessage, int, error) {
				log.WithField("method", method).Trace("jsonrpc: request")
				res, code, err := next(ctx, data)
				log.WithField("method", method).WithField("code", code).WithError(err).Trace("jsonrpc: response")
				return res, code, err
			}
		})
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	exitReceived := false
	exitFn := func() {
		exitReceived = true
		// Cancel first, then close the stream so a read blocked on a quiet
		// pipe wakes up and Serve can return.
		cancel()
		_ = conn.Close()
	}
	if err := s.registerMethods(rpc, exitFn); err != nil {
		return err
	}

	log.Info("Started laravel-ls server")
	err := jsonrpcstdio.Serve(ctx, rpc, jsonrpcstdio.FramingContentLength, conn, conn)

	switch {
	case err == nil, exitReceived && errors.Is(err, context.Canceled):
		// nil: the client closed our stdin. Canceled+exitReceived: the LSP
		// exit notification ended the session. Either way the connection is
		// down; per LSP it is an error unless shutdown was requested first.
		if !s.shutdownReceived {
			return fmt.Errorf("disconnected without an shutdown request")
		}
		return nil
	case errors.Is(err, context.Canceled):
		return fmt.Errorf("context closed")
	default:
		return err
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
