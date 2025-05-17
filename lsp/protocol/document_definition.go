package protocol

import (
	"encoding/json"
	"errors"
)

const (
	// MethodTextDocumentDefinition method name of "textDocument/definition".
	MethodTextDocumentDefinition = "textDocument/definition"
)

// DefinitionParams defines the parameters for a textDocument/definition request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#definitionParams
type DefinitionParams struct {
	WorkDoneProgressParams
	PartialResultParams
	TextDocumentPositionParams
}

// DefinitionResponse represents the result of a textDocument/definition request.
//
// It can be a single Location, a slice of Locations, a slice of LocationLinks or null.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#definition
type DefinitionResponse struct {
	Location      *Location
	LocationList  []Location
	LocationLinks []LocationLink
	Null          bool
}

func (dr DefinitionResponse) MarshalJSON() ([]byte, error) {
	if dr.Location != nil {
		return json.Marshal(dr.Location)
	}
	if dr.LocationList != nil {
		return json.Marshal(dr.LocationList)
	}
	if dr.LocationLinks != nil {
		return json.Marshal(dr.LocationLinks)
	}
	return []byte("null"), nil
}

func (dr *DefinitionResponse) UnmarshalJSON(data []byte) error {
	// Make sure object is reset.
	*dr = DefinitionResponse{}

	// Check for null
	if string(data) == "null" {
		dr.Null = true
		return nil
	}

	// Try single Location
	var loc Location
	if err := json.Unmarshal(data, &loc); err == nil && loc.URI != "" {
		dr.Location = &loc
		return nil
	}

	// Try []Location
	var locList []Location
	if err := json.Unmarshal(data, &locList); err == nil {
		dr.LocationList = locList
		return nil
	}

	// Try []LocationLink
	var linkList []LocationLink
	if err := json.Unmarshal(data, &linkList); err == nil {
		dr.LocationLinks = linkList
		return nil
	}

	return errors.New("invalid definition response: not null, Location, []Location, or []LocationLink")
}
