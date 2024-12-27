package protocol

const (
	MethodTextDocumentDidOpen = "textDocument/didOpen"

	MethodTextDocumentDidClose = "textDocument/didClose"

	MethodTextDocumentDidChange = "textDocument/didChange"

	MethodTextDocumentDidSave = "textDocument/didSave"
)

// DidOpenTextDocumentParams represents the parameters sent in the notification when a document is opened by the client.
type DidOpenTextDocumentParams struct {
	// TextDocument contains the details of the document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// DidCloseTextDocumentParams represents the parameters sent in the notification
// when a document is closed by the client.
type DidCloseTextDocumentParams struct {
	// TextDocument contains the details of the document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DidChangeTextDocumentParams represents the parameters sent in the notification
// when a document is changed by the client.
type DidChangeTextDocumentParams struct {
	// TextDocument holds the identifier and version of the text document that changed.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// ContentChanges is a list of changes applied to the document.
	// The changes represent the modified content of the document.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// DidSaveTextDocumentParams represents the parameters sent in the notification
// when a document is saved by the client.
type DidSaveTextDocumentParams struct {
	// TextDocument holds the identifier and version of the text document that was saved
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Text is the new text of the document that was saved.
	Text *string `json:"text,omitempty"`
}

// TextDocumentContentChangeEvent represents a change to a text document.
// It includes the text that was inserted, deleted, or modified.
type TextDocumentContentChangeEvent struct {
	// Range specifies the range of the text document that changed.
	// If nil, the entire document is considered changed.
	Range *Range `json:"range,omitempty"`

	// RangeLength is the length of the range that got replaced.
	// This field is optional and only provided when applicable.
	RangeLength *int `json:"rangeLength,omitempty"`

	// Text is the new text of the document after the change.
	Text string `json:"text"`
}
