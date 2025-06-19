package protocol

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

// The reason why code actions were requested.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeactiontriggerkind
//
// @since 3.17.0
type CodeActionTriggerKind int

const (
	// Code actions were explicitly requested by the user or by an extension.
	CodeActionTriggerKindInvoked CodeActionTriggerKind = 1

	// Code actions were requested automatically.
	// This typically happens when a diagnostic is reported and code actions are requested.
	// This kind of request should not trigger UI (e.g. prompt the user to pick a code action).
	CodeActionTriggerKindAutomatic CodeActionTriggerKind = 2
)

// CodeActionParams defines the parameters passed to the `textDocument/codeAction` request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeActionParams
type CodeActionParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range for which the command was invoked.
	Range Range `json:"range"`

	// Context carrying additional information.
	Context CodeActionContext `json:"context"`
}

// CodeActionContext contains additional diagnostic information about the context in which a code action is run.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeActionContext
type CodeActionContext struct {
	// An array of diagnostics known on the client side overlapping the range provided to the `textDocument/codeAction` request.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Requested kind of actions to return.
	// Actions not of this kind are filtered out by the client before being shown. So servers can omit computing them.
	// @since 3.15.0
	Only []CodeActionKind `json:"only,omitempty"`

	// The reason why code actions were requested.
	// @since 3.17.0
	TriggerKind CodeActionTriggerKind `json:"triggerKind,omitempty"`
}

// A code action represents a change that can be performed in code, e.g. to fix a problem or to refactor code.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeAction
type CodeAction struct {
	// A short, human-readable, title for this code action.
	Title string `json:"title"`

	// The kind of the code action.
	// Used to filter code actions.
	Kind CodeActionKind `json:"kind,omitempty"`

	// The diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	// Marks this as a preferred action.
	// Preferred actions are used by the `auto fix` command and can be targeted by keybindings.
	// @since 3.15.0
	IsPreferred bool `json:"isPreferred,omitempty"`

	// Marks that the code action cannot currently be applied.
	// Clients should follow the `disabled` property to determine if the action is shown in the UI.
	// @since 3.16.0
	Disabled *CodeActionDisabled `json:"disabled,omitempty"`

	// The workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	// A command this code action executes. If both `edit` and `command` are specified, edit is applied first.
	Command *Command `json:"command,omitempty"`

	// A data entry field that is preserved between a `textDocument/codeAction` request and a `codeAction/resolve` request.
	// @since 3.16.0
	Data any `json:"data,omitempty"`
}

// `CodeActionDisabled` is used to signal to the client that a code action is
// currently disabled and not applicable in the current context.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeactiondisabled
type CodeActionDisabled struct {
	// Human-readable description of why the code action is currently disabled.
	// This is displayed in the UI.
	Reason string `json:"reason"`
}
