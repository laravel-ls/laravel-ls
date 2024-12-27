package protocol

// ServerInfo represents additional information about the server.
type ServerInfo struct {
	// Name is the name of the server.
	Name string `json:"name"`

	// Version is the version of the server (optional).
	Version string `json:"version,omitempty"`
}
