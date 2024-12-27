package protocol

// Diagnostic represents a single diagnostic (error, warning, etc.) in a text document.
type Diagnostic struct {
	// Range specifies the location of the diagnostic within the document.
	Range Range `json:"range"`

	// Severity represents the severity of the diagnostic (e.g., error, warning).
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	// Code is an optional identifier for the diagnostic, which can be a string or number.
	Code interface{} `json:"code,omitempty"`

	// CodeDescription provides an optional URI pointing to more information about the diagnostic code.
	CodeDescription *CodeDescription `json:"codeDescription,omitempty"`

	// Source identifies the origin of this diagnostic (e.g., "typescript").
	Source string `json:"source,omitempty"`

	// Message is the diagnostic message to be displayed.
	Message string `json:"message"`

	// Tags is an optional array of tags for additional metadata.
	Tags []DiagnosticTag `json:"tags,omitempty"`

	// RelatedInformation is an optional array of additional locations related to this diagnostic.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`

	// Data provides additional metadata for the diagnostic.
	Data interface{} `json:"data,omitempty"`
}

// DiagnosticSeverity defines the severity level for a diagnostic.
type DiagnosticSeverity int

const (
	// SeverityError indicates an error.
	SeverityError DiagnosticSeverity = 1

	// SeverityWarning indicates a warning.
	SeverityWarning DiagnosticSeverity = 2

	// SeverityInformation indicates an informational message.
	SeverityInformation DiagnosticSeverity = 3

	// SeverityHint indicates a hint or suggestion.
	SeverityHint DiagnosticSeverity = 4
)

// DiagnosticTag represents additional metadata for diagnostics.
type DiagnosticTag int

const (
	// Unnecessary indicates that the diagnostic refers to code that is unnecessary or unused.
	Unnecessary DiagnosticTag = 1

	// Deprecated indicates that the diagnostic refers to deprecated code.
	Deprecated DiagnosticTag = 2
)

// CodeDescription provides a URI that points to additional information about a diagnostic code.
type CodeDescription struct {
	// Href is a URI pointing to more information about the diagnostic code.
	Href string `json:"href"`
}

// DiagnosticRelatedInformation represents additional diagnostic information related to the primary diagnostic.
type DiagnosticRelatedInformation struct {
	// Location is the location where this related diagnostic is reported.
	Location Location `json:"location"`

	// Message is the message to be displayed for this related diagnostic.
	Message string `json:"message"`
}
