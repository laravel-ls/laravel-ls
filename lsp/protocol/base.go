package protocol

// Position represents a position in a text document (line and character offset).
type Position struct {
	// Line is the zero-based line position in the document.
	Line int `json:"line"`

	// Character is the zero-based character offset in the line.
	Character int `json:"character"`
}

// Range represents a range in the text document, defined by a start and end position.
type Range struct {
	// Start is the start position of the range.
	Start Position `json:"start"`

	// End is the end position of the range.
	End Position `json:"end"`
}

// Language represents the programming language identifier for a document.
// It is used to specify the language type of the file (e.g., JavaScript, Python, etc.).
type LanguageID string

const (
	// LanguagePHP is the identifier for PHP documents.
	LanguagePHP LanguageID = "php"

	// LanguageBlade is the identifier for blade documents.
	LanguageBlade LanguageID = "blade"
)

// Location represents a specific location in a text document.
type Location struct {
	// URI is the unique resource identifier of the document.
	URI string `json:"uri"`

	// Range specifies the range within the document where the symbol's definition is located.
	Range Range `json:"range"`
}

// LocationLink provides additional metadata for a symbol's location, such as an origin selection range.
type LocationLink struct {
	// OriginSelectionRange is the range in the originating document where the request was initiated.
	// This can be used to specify the range of the symbol being requested.
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`

	// TargetURI is the URI of the target document where the definition is located.
	TargetURI string `json:"targetUri"`

	// TargetRange specifies the full range in the target document where the definition resides.
	TargetRange Range `json:"targetRange"`

	// TargetSelectionRange specifies the range within the target document that identifies the symbol definition.
	TargetSelectionRange Range `json:"targetSelectionRange"`
}
