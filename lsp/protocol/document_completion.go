package protocol

import (
	"encoding/json"
	"errors"
)

const (
	// MethodTextDocumentCompletion method name of "textDocument/completion".
	MethodTextDocumentCompletion = "textDocument/completion"
)

// CompletionTriggerKind specifies how the completion request was triggered.
type CompletionTriggerKind int

const (
	// Invoked means completion was manually triggered, such as by pressing Ctrl+Space.
	CompletionTriggerInvoked CompletionTriggerKind = 1

	// TriggerCharacter means completion was triggered by a specific character.
	CompletionTriggerCharacter CompletionTriggerKind = 2

	// TriggerForIncompleteCompletions means completion was triggered to resolve incomplete items.
	CompletionTriggerForIncompleteCompletions CompletionTriggerKind = 3
)

// CompletionContext provides additional information about the completion request context.
type CompletionContext struct {
	// TriggerKind specifies how the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	// TriggerCharacter is the character that triggered the completion, if applicable.
	// This field is only populated if the triggerKind is TriggerCharacter.
	TriggerCharacter *string `json:"triggerCharacter,omitempty"`
}

// CompletionParams represents the parameters for a `textDocument/completion` request.
// It provides the text document, position, context, and optional progress tracking.
type CompletionParams struct {
	// TextDocument identifies the text document where completion is being requested.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position specifies the position in the document where completion is requested.
	Position Position `json:"position"`

	// Context provides additional information about the context in which completion is requested.
	Context *CompletionContext `json:"context,omitempty"`

	// WorkDoneProgressParams allows reporting progress for this request.
	WorkDoneProgressParams
}

// CompletionResult represents the result of a `textDocument/completion` request.
// It can either be a CompletionList or a slice of CompletionItem.
type CompletionResult struct {
	// CompletionList is populated if the result is a CompletionList.
	CompletionList *CompletionList

	// CompletionItems is populated if the result is a slice of CompletionItem.
	CompletionItems []CompletionItem
}

// MarshalJSON customizes the JSON encoding for CompletionResult.
func (cr CompletionResult) MarshalJSON() ([]byte, error) {
	if cr.CompletionList != nil {
		return json.Marshal(cr.CompletionList)
	}
	return json.Marshal(cr.CompletionItems)
}

// UnmarshalJSON customizes the JSON decoding for CompletionResult.
func (cr *CompletionResult) UnmarshalJSON(data []byte) error {
	// Try to decode as a CompletionList first.
	var list CompletionList
	if err := json.Unmarshal(data, &list); err == nil && list.Items != nil {
		cr.CompletionList = &list
		return nil
	}

	// If decoding as CompletionList fails, try to decode as []CompletionItem.
	var items []CompletionItem
	if err := json.Unmarshal(data, &items); err == nil {
		cr.CompletionItems = items
		return nil
	}

	// If neither decoding attempt succeeds, return an error.
	return errors.New("invalid CompletionResult: data must be a CompletionList or []CompletionItem")
}
