package protocol

import (
	"encoding/json"
	"errors"
)

const (
	// MethodTextDocumentCodeAction method name of "textDocument/codeAction".
	MethodTextDocumentCodeAction = "textDocument/codeAction"
)

// CodeActionKind defines the type of code action.
type CodeActionKind string

const (
	CodeActionQuickFix              CodeActionKind = "quickfix"
	CodeActionRefactor              CodeActionKind = "refactor"
	CodeActionRefactorExtract       CodeActionKind = "refactor.extract"
	CodeActionRefactorInline        CodeActionKind = "refactor.inline"
	CodeActionRefactorRewrite       CodeActionKind = "refactor.rewrite"
	CodeActionSource                CodeActionKind = "source"
	CodeActionSourceOrganizeImports CodeActionKind = "source.organizeImports"
)

// CodeActionParams defines the parameters of a `textDocument/codeAction` request.
type CodeActionParams struct {
	// WorkDoneProgressParams are standard parameters for work progress reporting.
	WorkDoneProgressParams

	// PartialResultParams are standard parameters for partial result reporting.
	PartialResultParams

	// TextDocument specifies the text document in which code actions are requested.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Range specifies the range in the document where code actions are requested.
	Range Range `json:"range"`

	// Context provides additional information about the request, such as diagnostics.
	Context CodeActionContext `json:"context"`
}

// CodeActionContext contains additional information about the context in which a code action is requested.
type CodeActionContext struct {
	// Diagnostics contains the diagnostics for the specified range in the document.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Only is an optional list of code action kinds to return.
	// If specified, the server should only return these kinds.
	Only []CodeActionKind `json:"only,omitempty"`
}

// CodeAction represents a specific action that can be performed in a text document.
type CodeAction struct {
	// Title is a short, human-readable title for the code action.
	Title string `json:"title"`

	// Kind is the kind of the code action, such as "quickfix" or "refactor".
	// Used to filter and organize code actions.
	Kind CodeActionKind `json:"kind,omitempty"`

	// Diagnostics are the diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	// Edit specifies changes to be made to the workspace when this action is applied.
	// Mutually exclusive with Command.
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	// Command specifies a command to execute after applying the code action.
	// Mutually exclusive with Edit.
	Command *Command `json:"command,omitempty"`

	// Data is custom data the server may use to identify the code action.
	Data interface{} `json:"data,omitempty"`
}

// CodeActionResult represents the result of a `textDocument/codeAction` request.
// It can be either an array of Command or CodeAction, or null.
type CodeActionResult struct {
	Commands    []Command
	CodeActions []CodeAction
}

// UnmarshalJSON customizes the unmarshaling of CodeActionResult.
func (r *CodeActionResult) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		return nil
	}

	// Attempt to unmarshal as []CodeAction
	var actions []CodeAction
	if err := json.Unmarshal(data, &actions); err == nil {
		r.CodeActions = actions
		return nil
	}

	// Attempt to unmarshal as []Command
	var commands []Command
	if err := json.Unmarshal(data, &commands); err == nil {
		r.Commands = commands
		return nil
	}

	// If neither works, return an error
	return errors.New("invalid CodeActionResult: must be []Command, []CodeAction, or null")
}

// MarshalJSON customizes the marshaling of CodeActionResult.
func (r CodeActionResult) MarshalJSON() ([]byte, error) {
	if len(r.CodeActions) > 0 {
		return json.Marshal(r.CodeActions)
	}

	if len(r.Commands) > 0 {
		return json.Marshal(r.Commands)
	}

	// Default to null if empty
	return []byte("null"), nil
}
