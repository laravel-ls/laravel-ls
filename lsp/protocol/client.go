package protocol

// ClientInfo represents additional information about the client.
type ClientInfo struct {
	// Name is the name of the client.
	Name string `json:"name"`

	// Version is the version of the client (optional).
	Version string `json:"version,omitempty"`
}

// Command represents a command that can be executed in the client.
// Commands are provided as part of code actions, completion items, and other features.
type Command struct {
	// Title is a human-readable title for the command, displayed in the UI.
	Title string `json:"title"`

	// Command is the identifier of the actual command to execute.
	Command string `json:"command"`

	// Arguments are additional arguments that the command accepts.
	// These are specific to the command being executed.
	Arguments []interface{} `json:"arguments,omitempty"`
}
