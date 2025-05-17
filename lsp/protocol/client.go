package protocol

// Information about the client.
// The client provides this information during the initialize request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeParams
//
// @since 3.15.0
type ClientInfo struct {
	// The name of the client as defined by the client.
	// For example "vscode", "emacs" or "vim".
	Name string `json:"name"`

	// The client's version as defined by the client.
	Version *string `json:"version,omitempty"`
}

// Represents a reference to a command. Provides a title which will be used to represent a command in the UI
// and optionally a command identifier and arguments.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#command
type Command struct {
	// Title of the command, like `save`.
	Title string `json:"title"`

	// The identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command handler should be invoked with.
	Arguments []LSPAny `json:"arguments,omitempty"`
}
