package protocol

// Represents a diagnostic, such as a compiler error or warning.
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnostic
type Diagnostic struct {
	// The range at which the message applies.
	Range Range `json:"range"`

	// The diagnostic's severity. To avoid interpretation mismatches when a
	// server is used with different clients it is highly recommended that
	// servers always provide a severity value. If omitted, it’s recommended
	// for the client to interpret it as an Error severity.
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	// The diagnostic's code, which might appear in the user interface.
	Code any `json:"code,omitempty"` // integer or string

	// An optional property to describe the error code.
	//
	// @since 3.16.0
	CodeDescription *CodeDescription `json:"codeDescription,omitempty"`

	// A human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'.
	Source string `json:"source,omitempty"`

	// The diagnostic's message.
	Message string `json:"message"`

	// Additional metadata about the diagnostic.
	//
	// @since 3.15.0
	Tags []DiagnosticTag `json:"tags,omitempty"`

	// An array of related diagnostic information, e.g. when symbol-names within
	// a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`

	// A data entry field that is preserved between
	// a `textDocument/publishDiagnostics` notification and
	// `textDocument/codeAction` request.
	//
	// @since 3.16.0
	Data LSPAny `json:"data,omitempty"`
}

// The severity of a diagnostic message.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnosticSeverity
type DiagnosticSeverity int

const (
	// Reports an error.
	DiagnosticSeverityError DiagnosticSeverity = 1

	// Reports a warning.
	DiagnosticSeverityWarning DiagnosticSeverity = 2

	// Reports an information.
	DiagnosticSeverityInformation DiagnosticSeverity = 3

	// Reports a hint.
	DiagnosticSeverityHint DiagnosticSeverity = 4
)

// Additional metadata about a diagnostic.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnosticTag
//
// @since 3.15.0
type DiagnosticTag int

const (
	// Unused or unnecessary code.
	// Clients are allowed to render diagnostics with this tag faded out instead of having
	// an error squiggle.
	DiagnosticTagUnnecessary DiagnosticTag = 1

	// Deprecated or obsolete code.
	// Clients are allowed to rendered diagnostics with this tag strike-through.
	DiagnosticTagDeprecated DiagnosticTag = 2
)

// Structure to capture a description for an error code.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeDescription
//
// @since 3.16.0
type CodeDescription struct {
	// An URI to open with more information about the diagnostic error.
	Href DocumentURI `json:"href"`
}

// Represents a related message and source code location for a diagnostic.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnosticRelatedInformation
type DiagnosticRelatedInformation struct {
	// The location of this related diagnostic information.
	Location Location `json:"location"`

	// The message of this related diagnostic information.
	Message string `json:"message"`
}
