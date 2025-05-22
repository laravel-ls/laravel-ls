package protocol

import (
	"encoding/json"
	"fmt"
)

const (
	MethodTextDocumentDiagnostic = "textDocument/diagnostic"
)

// Parameters of the document diagnostic request.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentDiagnosticParams
type DocumentDiagnosticParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document to request diagnostics for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The additional identifier provided during registration.
	Identifier string `json:"identifier,omitempty"`

	// The current version of the document.
	// If provided, servers can avoid computing diagnostics again if the document version hasn’t changed.
	PreviousResultID string `json:"previousResultId,omitempty"`
}

// DocumentDiagnosticReport is either a full or an unchanged diagnostic report.
//
// @since 3.17.0
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentDiagnosticReport
type DocumentDiagnosticReport struct {
	Full      *FullDocumentDiagnosticReport
	Unchanged *UnchangedDocumentDiagnosticReport
}

// A full diagnostic report with a full set of problems.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#fullDocumentDiagnosticReport
type FullDocumentDiagnosticReport struct {
	Kind     string       `json:"kind"` // Should always be "full"
	ResultID string       `json:"resultId,omitempty"`
	Items    []Diagnostic `json:"items"`
}

// An unchanged diagnostic report indicating nothing has changed.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#unchangedDocumentDiagnosticReport
type UnchangedDocumentDiagnosticReport struct {
	Kind     string `json:"kind"` // Should always be "unchanged"
	ResultID string `json:"resultId"`
}

func (r *DocumentDiagnosticReport) UnmarshalJSON(data []byte) error {
	// reset object first.
	*r = DocumentDiagnosticReport{}

	var temp struct {
		Kind string `json:"kind"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp.Kind {
	case "full":
		var full FullDocumentDiagnosticReport
		if err := json.Unmarshal(data, &full); err != nil {
			return err
		}
		r.Full = &full
	case "unchanged":
		var unchanged UnchangedDocumentDiagnosticReport
		if err := json.Unmarshal(data, &unchanged); err != nil {
			return err
		}
		r.Unchanged = &unchanged
	default:
		return fmt.Errorf("unknown document diagnostic report kind: %s", temp.Kind)
	}

	return nil
}
