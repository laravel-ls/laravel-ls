package protocol

// ServerCapabilities defines the capabilities of the language server.
type ServerCapabilities struct {
	// TextDocumentSync defines how the server handles text document synchronization.
	//
	// If omitted it defaults to TextDocumentSyncKind.None`
	TextDocumentSync interface{} `json:"textDocumentSync,omitempty"` // *TextDocumentSyncOptions | TextDocumentSyncKind

	// HoverProvider indicates whether the server provides hover support.
	HoverProvider bool `json:"hoverProvider,omitempty"`

	// CompletionProvider defines the server's capability to provide completion items.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// SignatureHelpProvider defines the server's capability to provide signature help.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// DefinitionProvider represents the server's capability of providing definitions
	// for symbols in a text document, allowing the client to navigate to the symbol's definition.
	DefinitionProvider bool `json:"definitionProvider,omitempty"`

	// The server has support for pull model diagnostics.
	DiagnosticProvider interface{} `json:"diagnosticProvider,omitempty"`

	// The server has support for code actions.
	CodeActionProvider bool `json:"codeActionProvider,omitempty"`
}

// TextDocumentSyncKind defines how the host (editor) should sync document changes to the language server.
type TextDocumentSyncKind float64

const (
	// TextDocumentSyncKindNone documents should not be synced at all.
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	// TextDocumentSyncKindFull documents are synced by always sending the full content
	// of the document.
	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	// TextDocumentSyncKindIncremental documents are synced by sending the full content on open.
	// After that only incremental updates to the document are
	// send.
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

// CompletionOptions defines how the server supports providing completion items.
type CompletionOptions struct {
	// ResolveProvider indicates whether the server provides support to resolve additional information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	// TriggerCharacters defines the characters that trigger completion.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

// SignatureHelpOptions defines how the server supports providing signature help.
type SignatureHelpOptions struct {
	// TriggerCharacters defines the characters that trigger signature help.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

// DiagnosticOptions contains configuration options for the diagnostic provider.
type DiagnosticOptions struct {
	WorkDoneProgressOptions

	// Identifier is an optional string to identify this diagnostic request.
	Identifier string `json:"identifier,omitempty"`

	// InterFileDependencies indicates if the diagnostics depend on multiple files.
	InterFileDependencies bool `json:"interFileDependencies"`

	// WorkspaceDiagnostics indicates if workspace-wide diagnostics are supported.
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}

// WorkDoneProgressOptions indicates if a request supports work-done progress reporting.
type WorkDoneProgressOptions struct {
	// WorkDoneProgress is a flag that indicates whether progress reporting is supported.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
