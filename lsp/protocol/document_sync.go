package protocol

const (
	MethodTextDocumentDidOpen = "textDocument/didOpen"

	MethodTextDocumentDidClose = "textDocument/didClose"

	MethodTextDocumentDidChange = "textDocument/didChange"

	MethodTextDocumentDidSave = "textDocument/didSave"
)

// The parameters sent in a didOpen notification.
//
// @since 3.0.0

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didOpenTextDocumentParams
type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// The parameters sent in a didClose notification.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didCloseTextDocumentParams
type DidCloseTextDocumentParams struct {
	// The document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// The parameters sent in a didChange notification.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didChangeTextDocumentParams
type DidChangeTextDocumentParams struct {
	// The document that did change. The version number points
	// to the version after all provided content changes have been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// The actual content changes. The content changes describe single state
	// changes to the document. So if there are two content changes c1 (at
	// array index 0) and c2 (at array index 1) for a document in state S then
	// c1 moves the document from S to S' and c2 from S' to S''. So c1 is
	// computed on the state S and c2 is computed on the state S'.
	//
	// To mirror the content of a document using change events use the following
	// approach:
	// - start with the same initial content
	// - apply the 'textDocument/didChange' notifications in the order you
	//   receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single notification
	//   in the order you receive them.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// The parameters sent in a didSave notification.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didSaveTextDocumentParams
type DidSaveTextDocumentParams struct {
	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`
}

// An event describing a change to a text document.
// If range and rangeLength are omitted, the new text is considered the full content.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentContentChangeEvent
type TextDocumentContentChangeEvent struct {
	// The range of the document that changed.
	Range *Range `json:"range,omitempty"`

	// The optional length of the range that got replaced.
	// Deprecated: use range instead.
	RangeLength *uint `json:"rangeLength,omitempty"`

	// The new text of the document.
	Text string `json:"text"`
}
