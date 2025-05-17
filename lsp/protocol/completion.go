package protocol

// The kind of a completion entry.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemKind
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

// Tags are extra annotations that tweak the rendering of a completion item.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemTag
//
// @since 3.15.0
type CompletionItemTag int

const (
	// Render a completion as obsolete, usually using a strike-out.
	CompletionItemTagDeprecated CompletionItemTag = 1
)

// Defines how the insert text in a completion item is interpreted.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertTextFormat
type InsertTextFormat int

const (
	// The primary text to be inserted is treated as a plain string.
	InsertTextFormatPlainText InsertTextFormat = 1

	// The primary text to be inserted is treated as a snippet.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	InsertTextFormatSnippet InsertTextFormat = 2
)

// How whitespace and indentation is handled during completion item insertion.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertTextMode
//
// @since 3.16.0
type InsertTextMode int

const (
	// The insertion or replace strings is taken as it is. If the
	// value is multi line the lines below the cursor will be
	// inserted using the indentation defined in the string value.
	// The client will not apply any kind of adjustments to the
	// string.
	InsertTextModeAsIs InsertTextMode = 1

	// The editor adjusts leading whitespace of new lines so that
	// they match the indentation up to the cursor of the line for
	// which the item is accepted.
	//
	// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
	// multi line completion item is indented using 2 tabs and all
	// following lines inserted will be indented using 2 tabs as well.
	InsertTextModeAdjustIndentation InsertTextMode = 2
)

// A completion item represents a suggestion to complete text, typically
// during typing.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItem
type CompletionItem struct {
	// The label of this completion item.
	//
	// The label property is also by default the text that
	// is inserted when selecting this completion.
	//
	// If label details are provided the label itself should
	// be an unqualified name of the completion item.
	Label string `json:"label"`

	// Additional details for the label.
	//
	// @since 3.17.0
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`

	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor. The standardized set
	// of available values is defined in `CompletionItemKind`.
	Kind *CompletionItemKind `json:"kind,omitempty"`

	// Tags for this completion item.
	//
	// @since 3.15.0
	Tags []CompletionItemTag `json:"tags,omitempty"`

	// A human-readable string with additional information
	// about this item, like type or symbol information.
	Detail string `json:"detail,omitempty"`

	// A human-readable string that represents a doc-comment.
	Documentation string `json:"documentation,omitempty"`

	// Indicates if this item is deprecated.
	//
	// Deprecated: Use `tags` instead if supported.
	Deprecated *bool `json:"deprecated,omitempty"`

	// Select this item when showing.
	//
	// Note: that only one completion item can be selected and that the
	// tool / client decides which item that is. The rule is that the *first*
	// item of those that match best is selected.
	Preselect *bool `json:"preselect,omitempty"`

	// A string that should be used when comparing this item
	// with other items. When omitted the label is used
	// as the sort text for this item.
	SortText string `json:"sortText,omitempty"`

	// A string that should be used when filtering a set of
	// completion items. When omitted the label is used as the
	// filter text for this item.
	FilterText string `json:"filterText,omitempty"`

	// A string that should be inserted into a document when selecting
	// this completion. When omitted the label is used as the insert text
	// for this item.
	//
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally. For example
	// VS Code when code complete is requested in this example
	// `con<cursor position>` and a completion item with an `insertText` of
	// `console` is provided it will only insert `sole`. Therefore it is
	// recommended to use `textEdit` instead since it avoids additional client
	// side interpretation.
	InsertText string `json:"insertText,omitempty"`

	// The format of the insert text. The format applies to both the
	// `insertText` property and the `newText` property of a provided
	// `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
	//
	// Please note that the insertTextFormat doesn't apply to
	// `additionalTextEdits`.
	InsrtTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`

	// How whitespace and indentation is handled during completion
	// item insertion. If not provided the client's default value depends on
	// the `textDocument.completion.insertTextMode` client capability.
	//
	// @since 3.16.0
	// @since 3.17.0 - support for `textDocument.completion.insertTextMode`
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`

	// An edit which is applied to a document when selecting this completion.
	// When an edit is provided the value of `insertText` is ignored.
	//
	// *Note:* The range of the edit must be a single line range and it must
	// contain the position at which completion has been requested.
	//
	// Most editors support two different operations when accepting a completion
	// item. One is to insert a completion text and the other is to replace an
	// existing text with a completion text. Since this can usually not be
	// predetermined by a server it can report both ranges. Clients need to
	// signal support for `InsertReplaceEdit`s via the
	// `textDocument.completion.completionItem.insertReplaceSupport` client
	// capability property.
	//
	// Note 1: The text edit's range as well as both ranges from an insert
	// replace edit must be a [single line] and they must contain the position
	// at which completion has been requested.
	// Note 2: If an `InsertReplaceEdit` is returned the edit's insert range
	// must be a prefix of the edit's replace range, that means it must be
	// contained and starting at the same position.
	//
	// @since 3.16.0 additional type `InsertReplaceEdit`
	TextEdit *TextEdit `json:"textEdit,omitempty"`

	// The edit text used if the completion item is part of a CompletionList and
	// CompletionList defines an item default for the text edit range.
	//
	// Clients will only honor this property if they opt into completion list
	// item defaults using the capability `completionList.itemDefaults`.
	//
	// If not provided and a list's default range is provided the label
	// property is used as a text.
	//
	// @since 3.17.0
	TextEditText string `json:"textEditText,omitempty"`

	// An optional array of additional text edits that are applied when
	// selecting this completion. Edits must not overlap (including the same
	// insert position) with the main edit nor with themselves.
	//
	// Additional text edits should be used to change text unrelated to the
	// current cursor position (for example adding an import statement at the
	// top of the file if the completion item will insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	// An optional set of characters that when pressed while this completion is
	// active will accept it first and then type that character. *Note* that all
	// commit characters should have `length=1` and that superfluous characters
	// will be ignored.
	CommitCharacters string `json:"commitCharacters,omitempty"`

	// An optional command that is executed *after* inserting this completion.
	// Note that additional modifications to the current document should be
	// described with the additionalTextEdits-property.
	Command *Command `json:"command,omitempty"`

	// A data entry field that is preserved on a completion item between
	// a completion and a completion resolve request.
	Data LSPAny `json:"data,omitempty"`
}

// Additional details for a completion item label.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemLabelDetails
//
// @since 3.17.0
type CompletionItemLabelDetails struct {
	// An optional string which is rendered less prominently directly after `CompletionItem.label`, without any spacing.
	Detail *string `json:"detail,omitempty"`

	// An optional string which is rendered less prominently after `CompletionItem.detail`. It should be used for additional
	// information, like type or symbol information.
	Description *string `json:"description,omitempty"`
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
