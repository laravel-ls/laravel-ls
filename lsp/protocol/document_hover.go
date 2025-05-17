package protocol

import (
	"encoding/json"
	"errors"
)

const (
	MethodTextDocumentHover = "textDocument/hover"
)

// Parameters for a textDocument/hover request.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#hoverParams
type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// The result of a hover request.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#hover
type Hover struct {
	// The hover's content.
	Contents MarkupContentOrMarkedString `json:"contents"`

	// An optional range inside the text document that is used to
	// visualize the hover, e.g. by changing the background color.
	Range *Range `json:"range,omitempty"`
}

// The result of a textDocument/hover request.
//
// Can be a Hover object or null.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#hover
type HoverResult struct {
	Hover *Hover
	Null  bool
}

func (h *HoverResult) UnmarshalJSON(data []byte) error {
	*h = HoverResult{}

	if string(data) == "null" {
		h.Null = true
		return nil
	}
	var hover Hover
	if err := json.Unmarshal(data, &hover); err == nil {
		return err
	}
	h.Hover = &hover
	return nil
}

func (h HoverResult) MarshalJSON() ([]byte, error) {
	if h.Null {
		return []byte("null"), nil
	}
	return json.Marshal(h.Hover)
}

// MarkedString can be used to render human-readable text,
// optionally with a language hint.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#markedString
type MarkedString struct {
	Language string `json:"language,omitempty"`
	Value    string `json:"value"`
}

// MarkupContent represents a string value with optional markup kind.
//
// @since 3.0.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#markupcontent
type MarkupContent struct {
	Kind  string `json:"kind"`  // "plaintext" or "markdown"
	Value string `json:"value"` // actual content
}

type MarkupContentOrMarkedString struct {
	Markup        *MarkupContent
	MarkedString  *MarkedString
	MarkedStrings []MarkedString
}

func (m *MarkupContentOrMarkedString) UnmarshalJSON(data []byte) error {
	var markup MarkupContent
	if err := json.Unmarshal(data, &markup); err == nil && markup.Kind != "" {
		m.Markup = &markup
		return nil
	}

	var single MarkedString
	if err := json.Unmarshal(data, &single); err == nil && (single.Language != "" || single.Value != "") {
		m.MarkedString = &single
		return nil
	}

	var many []MarkedString
	if err := json.Unmarshal(data, &many); err == nil {
		m.MarkedStrings = many
		return nil
	}

	return errors.New("invalid contents: not MarkupContent, MarkedString, or []MarkedString")
}
