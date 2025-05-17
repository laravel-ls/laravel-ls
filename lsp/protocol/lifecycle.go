package protocol

const (
	MethodInitialize = "initialize"

	MethodInitialized = "initialized"
)

// The initialize parameters.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeParams
type InitializeParams struct {
	WorkDoneProgressParams

	// The process Id of the parent process that started the server.
	// Is null if the process has not been started by another process.
	ProcessID *int `json:"processId,omitempty"`

	// The information about the client.
	//
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// The locale the client is currently showing the user interface in.
	// This must not be used to influence the content of hover, completion, signature help etc.
	//
	// @since 3.16.0
	Locale *string `json:"locale,omitempty"`

	// The rootPath of the workspace. Is null if no folder is open.
	//
	// @deprecated in favour of rootUri.
	RootPath string `json:"rootPath,omitempty"`

	// The rootUri of the workspace. Is null if no folder is open.
	// If both `rootPath` and `rootUri` are set `rootUri` wins.
	RootURI DocumentURI `json:"rootUri,omitempty"`

	// User provided initialization options.
	InitializationOptions LSPAny `json:"initializationOptions,omitempty"`

	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`

	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitempty"`

	// The workspace folders configured in the client when the server starts.
	// This property is only available if the client supports workspace folders.
	//
	// @since 3.6.0
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// The result returned from an initialize request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
type InitializeResult struct {
	// The capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// Information about the server.
	//
	// @since 3.15.0
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`
}

// The LSP allows the client to control the tracing of the server.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#traceValue
type TraceValue string

const (
	// No tracing.
	TraceValueOff TraceValue = "off"

	// Messages only.
	TraceValueMessages TraceValue = "messages"

	// Verbose message logging.
	TraceValueVerbose TraceValue = "verbose"
)
