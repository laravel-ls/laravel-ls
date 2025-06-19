package protocol

// DocumentChangeOperation represents any valid document change operation.
type DocumentChangeOperation interface {
	isDocumentChangeOperation()
}

// A generic resource operation.
type ResourceOperation struct {
	// resource operation kind
	Kind string `json:"kind"`

	// An optional annotation identifier describing the operation.
	AnnotationID string `json:"annotationId,omitempty"`
}

// CreateFile operation.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#createFile
type CreateFile struct {
	ResourceOperation
	URI     DocumentURI        `json:"uri"`
	Options *CreateFileOptions `json:"options,omitempty"`
}

func (CreateFile) isDocumentChangeOperation() {}

// Options to create a file.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#createFileOptions
type CreateFileOptions struct {
	// Overwrite existing file. Overwrite wins over ignoreIfExists.
	Overwrite bool `json:"overwrite,omitempty"`

	// Ignore if exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// RenameFile operation.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#renameFile
type RenameFile struct {
	ResourceOperation
	OldURI  DocumentURI        `json:"oldUri"`
	NewURI  DocumentURI        `json:"newUri"`
	Options *RenameFileOptions `json:"options,omitempty"`
}

// Options to rename a file.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#renameFileOptions
type RenameFileOptions struct {
	// Overwrite target if existing.
	Overwrite bool `json:"overwrite,omitempty"`

	// Ignores if target exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

func (RenameFile) isDocumentChangeOperation() {}

// DeleteFile operation.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#deleteFile
type DeleteFile struct {
	ResourceOperation
	URI     DocumentURI        `json:"uri"`
	Options *DeleteFileOptions `json:"options,omitempty"`
}

func (DeleteFile) isDocumentChangeOperation() {}

// Options to delete a file.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#deleteFileOptions
type DeleteFileOptions struct {
	// Delete content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`

	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}
