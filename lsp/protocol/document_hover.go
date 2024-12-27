package protocol

import (
	"encoding/json"
)

const (
	MethodTextDocumentHover = "textDocument/hover"
)

// interface HoverParams {
// 	textDocument: string; /** The text document's URI in string form */
// 	position: { line: uinteger; character: uinteger; };
// }

// HoverParams represents the parameters for a `textDocument/hover` request.
// It provides the position in the text document where hover information is requested.
type HoverParams struct {
	// TextDocument is the identifier for the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position specifies the position inside the text document where hover information is requested.
	Position Position `json:"position"`

	// WorkDoneProgressParams allows reporting progress for this request.
	WorkDoneProgressParams
}

// HoverResult represents the result of a `textDocument/hover` request.
// It contains the hover content and an optional range within the document.
type HoverResult struct {
	// Contents is the hover information provided by the server.
	// It can be a string, a MarkupContent object, or an array of MarkedString objects.
	Contents HoverContent `json:"contents"`

	// Range is an optional range in the document that this hover refers to.
	Range *Range `json:"range,omitempty"`
}

// MarkedString represents either a raw string or a language-specific string with a language identifier.
type MarkedString struct {
	Language string `json:"language,omitempty"`
	Value    string `json:"value"`
}

// MarkupContent represents content with a specific markup kind (e.g., Markdown or plain text).
type MarkupContent struct {
	// Kind specifies the type of markup (e.g., "plaintext" or "markdown").
	Kind string `json:"kind"`

	// Value is the content to be displayed in the hover.
	Value string `json:"value"`
}

// HoverContent represents the content for the hover response.
// It can be one of MarkupContent, MarkedString, or an array of MarkedString.
type HoverContent struct {
	MarkupContent   *MarkupContent `json:"-"`
	MarkedStrings   []MarkedString `json:"-"`
	PlainTextString *string        `json:"-"`
}

// MarshalJSON customizes JSON encoding for HoverContent.
func (h HoverContent) MarshalJSON() ([]byte, error) {
	if h.MarkupContent != nil {
		return json.Marshal(h.MarkupContent)
	}
	if h.MarkedStrings != nil {
		return json.Marshal(h.MarkedStrings)
	}
	if h.PlainTextString != nil {
		return json.Marshal(h.PlainTextString)
	}
	return nil, nil
}
