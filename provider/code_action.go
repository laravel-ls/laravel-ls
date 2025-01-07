package provider

import (
	"github.com/laravel-ls/laravel-ls/lsp/protocol"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type CodeActionPublish func(protocol.CodeAction)

type CodeActionContext struct {
	BaseContext

	Range ts.Range

	Publish CodeActionPublish
}

// Interface that providers that supports code actions can implement.
type CodeActionProvider interface {
	ResolveCodeAction(CodeActionContext)
}
