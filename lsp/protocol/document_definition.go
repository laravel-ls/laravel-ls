package protocol

import "encoding/json"

const (
	// MethodTextDocumentDefinition method name of "textDocument/definition".
	MethodTextDocumentDefinition = "textDocument/definition"
)

// DefinitionParams represents the parameters for the `textDocument/definition` request.
// This request is used to find the definition of a symbol at a specific position in a text document.
type DefinitionParams struct {
	// WorkDoneProgressParams allows the client to request progress updates for the operation.
	WorkDoneProgressParams

	// PartialResultParams allows the server to send partial results for the request.
	PartialResultParams

	// TextDocument specifies the document where the symbol is located.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position specifies the position within the document where the symbol is located.
	Position Position `json:"position"`
}

// DefinitionResponse represents the response for the `textDocument/definition` request.
// It can be a single Location, a slice of Locations, or a slice of LocationLinks.
type DefinitionResponse struct {
	Location      *Location      // Single Location
	Locations     []Location     // Multiple Locations
	LocationLinks []LocationLink // Multiple LocationLinks
}

// MarshalJSON customizes the JSON encoding for DefinitionResponse.
func (r DefinitionResponse) MarshalJSON() ([]byte, error) {
	// If only a single Location is populated
	if r.Location != nil {
		return json.Marshal(r.Location)
	}

	// If multiple Locations are populated
	if len(r.Locations) > 0 {
		return json.Marshal(r.Locations)
	}

	// If multiple LocationLinks are populated
	if len(r.LocationLinks) > 0 {
		return json.Marshal(r.LocationLinks)
	}

	// If none are populated, return "null"
	return []byte("null"), nil
}
