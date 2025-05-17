package protocol

import (
	"encoding/json"
	"errors"
)

const (
	// MethodTextDocumentCompletion method name of "textDocument/completion".
	MethodTextDocumentCompletion = "textDocument/completion"
)

// How a completion was triggered.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionTriggerKind
type CompletionTriggerKind int

const (
	// Completion was triggered by typing an identifier (24x7 code complete), manual invocation (e.g Ctrl+Space) or via API.
	CompletionTriggerKindInvoked CompletionTriggerKind = 1

	// Completion was triggered by a trigger character specified by the `triggerCharacters` properties of the `CompletionOptions`.
	CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2

	// Completion was re-triggered as the current completion list is incomplete.
	CompletionTriggerKindTriggerForIncompleteCompletions CompletionTriggerKind = 3
)

// Contains additional information about the context in which a completion request is triggered.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionContext
type CompletionContext struct {
	// How the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	// The trigger character (a single character) that has trigger code complete.
	// Is undefined if `triggerKind !== CompletionTriggerKindTriggerCharacter`
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
}

// Completion parameters.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionParams
type CompletionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	// The completion context. This is only available if the client specifies to send this using the client capability `textDocument.completion.contextSupport === true`
	Context *CompletionContext `json:"context,omitempty"`
}

// The result of a textDocument/completion request is either an array of CompletionItem
// or a CompletionList.

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completion
type CompletionResponse struct {
	// When the result is a list of completion items.
	Items []CompletionItem

	// When the result is a completion list with additional metadata like isIncomplete.
	List *CompletionList

	// If null
	Null bool
}

// MarshalJSON customizes the JSON encoding for CompletionResponse.
func (cr CompletionResponse) MarshalJSON() ([]byte, error) {
	if cr.List != nil {
		return json.Marshal(cr.List)
	}
	return json.Marshal(cr.Items)
}

// UnmarshalJSON customizes the JSON decoding for CompletionResponse.
func (cr *CompletionResponse) UnmarshalJSON(data []byte) error {
	// Check for null
	if string(data) == "null" {
		cr.Null = true
		cr.List = nil
		cr.Items = nil
		return nil
	}

	// Try to decode as a CompletionList.
	var list CompletionList
	if err := json.Unmarshal(data, &list); err == nil && list.Items != nil {
		cr.Null = true
		cr.List = &list
		cr.Items = nil
		return nil
	}

	// If decoding as CompletionList fails, try to decode as []CompletionItem.
	var items []CompletionItem
	if err := json.Unmarshal(data, &items); err == nil {
		cr.Null = true
		cr.List = nil
		cr.Items = items
		return nil
	}

	// Unknown structure
	return errors.New("invalid CompletionResponse: not a CompletionList, []CompletionItem or null")
}
