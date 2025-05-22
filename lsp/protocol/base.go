package protocol

// The LSP any type
// It can be a primitive (string, number, boolean, null), an object, or an array.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#lspAny
//
// @since 3.17.0
type LSPAny = any

// A URI for a document. This is just a string alias, typically a `file://` or other scheme URI.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentUri
type DocumentURI string

// Position in a text document expressed as zero-based line and character offset.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#position
type Position struct {
	// Line position in a document (zero-based).
	Line uint32 `json:"line"`

	// Character offset on a line in a document (zero-based).
	//
	// The meaning of this offset is determined by the negotiated
	// `PositionEncodingKind`.
	Character uint32 `json:"character"`
}

// A set of predefined position encoding kinds.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#positionEncodingKind
//
// @since 3.17.0
type PositionEncodingKind string

const (
	// Character offsets count UTF-8 code units.
	PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"

	// Character offsets count UTF-16 code units.
	// This is the default and must always be supported by servers
	PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"

	// Character offsets count UTF-32 code units.
	// Implementation note: these are the same as Unicode code points,
	// so this `PositionEncodingKind` may also be used for an
	// encoding-agnostic representation of character offsets.
	PositionEncodingKindUTF32 PositionEncodingKind = "utf-32"
)

// A range in a text document expressed as (zero-based) start and end positions.
// A range is comparable to a selection in an editor. Therefore, the end position is exclusive.
// If you want to specify a range that contains a line including the line ending character(s)
// then use an end position denoting the start of the next line.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#range
type Range struct {
	// The range's start position.
	Start Position `json:"start"`

	// The range's end position.
	End Position `json:"end"`
}

// A language id represented as a string.
type LanguageID string

const (
	// LanguagePHP is the identifier for PHP documents.
	LanguagePHP LanguageID = "php"

	// LanguageBlade is the identifier for blade documents.
	LanguageBlade LanguageID = "blade"
)

// Represents a location inside a resource, such as a line inside a text file.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#location
type Location struct {
	// URI of the document.
	URI DocumentURI `json:"uri"`

	// The range within the document.
	Range Range `json:"range"`
}

// Represents a link between a source and a target location.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#locationLink
//
// @since 3.14.0
type LocationLink struct {
	// Span of the origin of this link.
	// Used as the underlined span for mouse interaction.
	// Clients should omit this property if the origin selection range is not applicable.
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`

	// The target resource identifier.
	TargetURI DocumentURI `json:"targetUri"`

	// The full target range of this link.
	// If the target for example is a symbol then target range is the range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"targetRange"`

	// The range that should be selected and revealed when this link is being followed, e.g. the name of a function.
	// Must be contained by the `targetRange`. See also `DocumentSymbol#range`.
	TargetSelectionRange Range `json:"targetSelectionRange"`
}

// A change annotation represents additional information that describes a change.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#changeAnnotation
//
// @since 3.16.0
type ChangeAnnotation struct {
	// A human-readable string describing the actual change. The string is
	// rendered prominent in the user interface.
	Label string `json:"label"`

	// A flag which indicates that user confirmation is needed
	// before applying the change.
	NeedsConfirmation *bool `json:"needsConfirmation,omitempty"`

	// A human-readable string which is rendered less prominent in the
	// user interface.
	Description string `json:"description,omitempty"`
}

// Describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
//
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
//
// @see https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#markupContent
type MarkupKind string

const (
	// Plain text is supported as a content format
	MarkupKindPlainText MarkupKind = "plaintext"
	// Markdown is supported as a content format
	MarkupKindMarkdown MarkupKind = "markdown"
)
