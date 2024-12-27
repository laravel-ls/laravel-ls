package protocol

const (
	MethodInitialize = "initialize"

	MethodInitialized = "initialized"
)

// InitializeParams represents the parameters sent during the "initialize" request.
type InitializeParams struct {
	WorkDoneProgressParams

	// ProcessID is the ID of the parent process that started the server. If null, the process is not specified.
	ProcessID int `json:"processId,omitempty"`

	// ClientInfo is the information about the client.
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// RootURI is the root URI of the workspace, which may be null.
	RootURI string `json:"rootUri,omitempty"`

	// Capabilities represents the client capabilities.
	Capabilities ClientCapabilities `json:"capabilities"`

	// Trace enables tracing and can be set to "off", "messages", or "verbose".
	Trace TraceValue `json:"trace,omitempty"`
}

// InitializeResult represents the response sent by the server after receiving
// an InitializeRequest. It contains the capabilities of the language server.
type InitializeResult struct {
	// Capabilities defines the capabilities provided by the language server.
	Capabilities ServerCapabilities `json:"capabilities"`

	// ServerInfo provides information about the server (optional).
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`
}

// TraceValue represents a InitializeParams Trace mode.
type TraceValue string

// list of TraceValue.
const (
	// TraceOff disable tracing.
	TraceOff TraceValue = "off"

	// TraceMessage normal tracing mode.
	TraceMessage TraceValue = "message"

	// TraceVerbose verbose tracing mode.
	TraceVerbose TraceValue = "verbose"
)
