package protocol

// The server info returned from an initialize request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
//
// @since 3.15.0
type ServerInfo struct {
	// The name of the server as defined by the server.
	Name string `json:"name"`

	// The server's version, if provided.
	Version string `json:"version,omitempty"`
}
