package protocol

// WorkspaceEdit represents changes to many resources in the workspace.
type WorkspaceEdit struct {
	// Changes is a map of document URIs to an array of TextEdits.
	// Changes should be applied to the documents in the specified order.
	Changes map[string][]TextEdit `json:"changes,omitempty"`

	// DocumentChanges is an optional array of TextDocumentEdits.
	// These edits support versioned documents and resource operations.
	DocumentChanges []TextDocumentEdit `json:"documentChanges,omitempty"`
}
