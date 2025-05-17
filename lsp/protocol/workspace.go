package protocol

// A workspace edit represents changes to many resources managed in the workspace.
// The edit should either provide `changes` or `documentChanges`. If the client can handle versioned document edits
// and if `documentChanges` are present, the latter are preferred over `changes`.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceEdit
type WorkspaceEdit struct {
	// Holds changes to existing resources.
	// Each entry describes edits to a single document.
	Changes map[DocumentURI][]TextEdit `json:"changes,omitempty"`

	// Depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes are either
	// an array of `TextDocumentEdit`s or a mix of `TextDocumentEdit`, `CreateFile`, `RenameFile`, and `DeleteFile` operations.
	//
	// Changes to existing files are described with `TextDocumentEdit`.
	// Changes to the workspace are described with resource operations like `RenameFile`, `CreateFile`, and `DeleteFile`.
	//
	// @since 3.13.0
	DocumentChanges []LSPAny `json:"documentChanges,omitempty"`

	// A map of change annotations that can be referenced in `AnnotatedTextEdit`s or create, rename and delete file operations.
	//
	// @since 3.16.0
	ChangeAnnotations map[string]ChangeAnnotation `json:"changeAnnotations,omitempty"`
}

// A workspace folder inside a client.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceFolder
//
// @since 3.6.0
type WorkspaceFolder struct {
	// The associated URI for this workspace folder.
	URI DocumentURI `json:"uri"`

	// The name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string `json:"name"`
}
