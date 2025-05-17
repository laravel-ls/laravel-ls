package protocol

// ServerCapabilities defines the capabilities of the language server.
type ServerCapabilities struct {
	// The position encoding the server picked from the encodings offered
	// by the client via the client capability `general.positionEncodings`.
	//
	// If the client didn't provide any position encodings the only valid
	// value that a server can return is 'utf-16'.
	//
	// If omitted it defaults to 'utf-16'.
	//
	// @since 3.17.0
	PositionEncoding *PositionEncodingKind `json:"positionEncoding,omitempty"`

	// Defines how text documents are synced. Is either a detailed structure
	// defining each notification or for backwards compatibility the
	// TextDocumentSyncKind number.
	// If omitted it defaults to `TextDocumentSyncKind.None`.
	TextDocumentSync any `json:"textDocumentSync,omitempty"` // *TextDocumentSyncOptions | TextDocumentSyncKind

	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// The server provides hover support.
	HoverProvider bool `json:"hoverProvider,omitempty"`

	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// The server provides goto definition support.
	DefinitionProvider bool `json:"definitionProvider,omitempty"`

	// The server provides code actions. The `CodeActionOptions` return type is
	// only valid if the client signals code action literal support via the
	// property `textDocument.codeAction.codeActionLiteralSupport`.
	CodeActionProvider bool `json:"codeActionProvider,omitempty"`

	// The server has support for pull model diagnostics.
	//
	// @since 3.17.0
	DiagnosticProvider any `json:"diagnosticProvider,omitempty"` // DiagnosticOptions | DiagnosticRegistrationOptions

	// Experimental server capabilities.
	Experimental *LSPAny `json:"experimental,omitempty"`
}

// Defines how the host (editor) should sync document changes to the language server.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentSyncKind
type TextDocumentSyncKind int

const (
	// Documents should not be synced at all.
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	// Documents are synced by always sending the full content of the document.
	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	// Documents are synced by sending the full content on open and close.
	// After that only incremental updates to the document are sent.
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

// Completion options used during `initialize` or server capabilities registration.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionOptions
type CompletionOptions struct {
	WorkDoneProgressOptions

	// The additional characters, beyond the defaults provided by the client (typically
	// [a-zA-Z]), that should automatically trigger a completion request. For example
	// `.` in JavaScript represents the beginning of an object property or method and is
	// thus a good candidate for triggering a completion request.
	//
	// Most tools trigger a completion request automatically without explicitly
	// requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
	// do so when the user starts to type an identifier. For example if the user
	// types `c` in a JavaScript file code complete will automatically pop up
	// present `console` besides others as a completion item. Characters that
	// make up identifiers don't need to be listed here.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// The list of all possible characters that commit a completion. This field
	// can be used if clients don't support individual commit characters per
	// completion item. See client capability
	// `completion.completionItem.commitCharactersSupport`.
	//
	// If a server provides both `allCommitCharacters` and commit characters on
	// an individual completion item the ones on the completion item win.
	//
	// @since 3.2.0
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`

	// The server supports the following `CompletionItem` specific
	// capabilities.
	//
	// @since 3.17.0
	CompletionItem *CompletionItemCapability `json:"completionItem,omitempty"`
}

// Capabilities specific to the `CompletionItem`.
//
// @since 3.17.0
type CompletionItemCapability struct {
	// The server has support for completion item label details (see also `CompletionItemLabelDetails`)
	// when receiving a completion item in a resolve call.
	// @since 3.17.0
	LabelDetailsSupport *bool `json:"labelDetailsSupport,omitempty"`
}

// Server capabilities for signature help requests.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelpOptions
type SignatureHelpOptions struct {
	// The characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// List of characters that re-trigger signature help.
	// These are typically the same as `triggerCharacters`, but may contain additional characters.
	// @since 3.2.0
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// Diagnostic registration options to configure pull diagnostics.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnosticOptions
//
// @since 3.17.0
type DiagnosticOptions struct {
	WorkDoneProgressOptions

	// An optional identifier under which the diagnostics are
	// managed by the client.
	Identifier string `json:"identifier,omitempty"`

	// Whether the language has inter file dependencies meaning that
	// editing code in one file can result in a different diagnostic
	// set in another file. Inter file dependencies are common for
	// most programming languages and typically uncommon for linters.
	InterFileDependencies bool `json:"interFileDependencies"`

	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}

// WorkDoneProgressOptions indicates if a request supports work-done progress reporting.
type WorkDoneProgressOptions struct {
	// WorkDoneProgress is a flag that indicates whether progress reporting is supported.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
