package protocol

// CompletionItemKind defines the type of completion item (e.g., Function, Class).
type CompletionItemKind int

const (
	CompletionItemKindNone          CompletionItemKind = 0
	CompletionItemKindText          CompletionItemKind = 1
	CompletionItemKindMethod        CompletionItemKind = 2
	CompletionItemKindFunction      CompletionItemKind = 3
	CompletionItemKindConstructor   CompletionItemKind = 4
	CompletionItemKindField         CompletionItemKind = 5
	CompletionItemKindVariable      CompletionItemKind = 6
	CompletionItemKindClass         CompletionItemKind = 7
	CompletionItemKindInterface     CompletionItemKind = 8
	CompletionItemKindModule        CompletionItemKind = 9
	CompletionItemKindProperty      CompletionItemKind = 10
	CompletionItemKindUnit          CompletionItemKind = 11
	CompletionItemKindValue         CompletionItemKind = 12
	CompletionItemKindEnum          CompletionItemKind = 13
	CompletionItemKindKeyword       CompletionItemKind = 14
	CompletionItemKindSnippet       CompletionItemKind = 15
	CompletionItemKindColor         CompletionItemKind = 16
	CompletionItemKindFile          CompletionItemKind = 17
	CompletionItemKindReference     CompletionItemKind = 18
	CompletionItemKindFolder        CompletionItemKind = 19
	CompletionItemKindEnumMember    CompletionItemKind = 20
	CompletionItemKindConstant      CompletionItemKind = 21
	CompletionItemKindStruct        CompletionItemKind = 22
	CompletionItemKindEvent         CompletionItemKind = 23
	CompletionItemKindOperator      CompletionItemKind = 24
	CompletionItemKindTypeParameter CompletionItemKind = 25
)

// CompletionItemTag defines optional tags for CompletionItem.
type CompletionItemTag int

const (
	CompletionItemTagDeprecated CompletionItemTag = 1
)

// InsertTextFormat specifies how InsertText should be interpreted.
type InsertTextFormat int

const (
	InsertTextFormatPlainText InsertTextFormat = 1 // Plain text format
	InsertTextFormatSnippet   InsertTextFormat = 2 // Snippet format
)

// InsertTextMode defines how whitespace is handled during insertion.
type InsertTextMode int

const (
	InsertTextModeAsIs              InsertTextMode = 1
	InsertTextModeAdjustIndentation InsertTextMode = 2
)

// CompletionItem represents a single completion suggestion.
type CompletionItem struct {
	// Label is the label of this completion item. It will be displayed in the UI.
	Label string `json:"label"`

	// Kind specifies the type of completion item (e.g., Function, Class).
	Kind CompletionItemKind `json:"kind,omitempty"`

	// Tags are optional tags for this item (e.g., Deprecated).
	Tags []CompletionItemTag `json:"tags,omitempty"`

	// Detail provides additional information about this item, like the type or symbol signature.
	Detail string `json:"detail,omitempty"`

	// Documentation contains a human-readable description of the item.
	Documentation *MarkupContent `json:"documentation,omitempty"`

	// Deprecated indicates whether this item is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`

	// Preselect determines if this item should be preselected in the UI.
	Preselect *bool `json:"preselect,omitempty"`

	// SortText is used to order completion items in the UI.
	SortText string `json:"sortText,omitempty"`

	// FilterText is used to filter items when performing fuzzy matching.
	FilterText string `json:"filterText,omitempty"`

	// InsertText is the string to insert when this item is selected.
	InsertText string `json:"insertText,omitempty"`

	// InsertTextFormat specifies how `InsertText` should be interpreted.
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`

	// InsertTextMode specifies how whitespace and indentation should be handled during insertion.
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`

	// TextEdit specifies changes to the document that occur when this item is selected.
	TextEdit *TextEdit `json:"textEdit,omitempty"`

	// AdditionalTextEdits specifies additional edits that are applied after this item is selected.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	// CommitCharacters is a list of characters that commit the completion item selection.
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// Command is an optional command that is executed after selecting this item.
	Command *Command `json:"command,omitempty"`

	// Data is a custom data field, which can be used by the server to identify or resolve this item.
	Data interface{} `json:"data,omitempty"`
}

// CompletionList represents a collection of completion items.
// It can be either a list of items or a flag indicating if further items can be resolved.
type CompletionList struct {
	// IsIncomplete indicates if the list is complete.
	// If true, the client should re-trigger completion when typing more characters.
	IsIncomplete bool `json:"isIncomplete"`

	// Items contains the completion items.
	Items []CompletionItem `json:"items"`
}
