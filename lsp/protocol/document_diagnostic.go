package protocol

const (
	MethodTextDocumentDiagnostic = "textDocument/diagnostic"
)

// DocumentDiagnosticParams represents the parameters for the `textDocument/diagnostic` request.
// It is used to request diagnostics for a specific document.
type DocumentDiagnosticParams struct {
	// WorkDoneProgressParams contains options for work done progress reporting.
	WorkDoneProgressParams

	// TextDocument identifies the specific document to retrieve diagnostics for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Identifier uniquely identifies the request.
	Identifier string `json:"identifier,omitempty"`

	// PreviousResultId is the identifier of the last known result for this document.
	// The server may use this to optimize diagnostic calculations and only return differences.
	PreviousResultId string `json:"previousResultId,omitempty"`
}

// DocumentDiagnosticReport is an interface for diagnostic reports,
// which can be either FullDocumentDiagnosticReport or RelatedDocumentDiagnosticReport.
type DocumentDiagnosticReport interface {
	GetKind() string
}

// FullDocumentDiagnosticReport represents a full set of diagnostics for a document.
type FullDocumentDiagnosticReport struct {
	// Kind indicates this is a full document diagnostic report.
	Kind string `json:"kind"` // Should be "full"

	// ResultId is an optional identifier for caching and incremental updates.
	ResultId string `json:"resultId,omitempty"`

	// Items contains the list of diagnostics for the document.
	Items []Diagnostic `json:"items"`
}

// RelatedDocumentDiagnosticReport represents diagnostics for a document with references to related documents.
type RelatedDocumentDiagnosticReport struct {
	// Kind indicates this is a related document diagnostic report.
	Kind string `json:"kind"` // Should be "related"

	// RelatedDocuments maps document URIs to associated DocumentDiagnosticReport objects.
	RelatedDocuments map[string]DocumentDiagnosticReport `json:"relatedDocuments"`
}

// GetKind returns the kind of FullDocumentDiagnosticReport.
func (f *FullDocumentDiagnosticReport) GetKind() string {
	return f.Kind
}

// GetKind returns the kind of RelatedDocumentDiagnosticReport.
func (r *RelatedDocumentDiagnosticReport) GetKind() string {
	return r.Kind
}
