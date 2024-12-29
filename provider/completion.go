package provider

import (
	"laravel-ls/lsp/protocol"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type CompletionPublish func(protocol.CompletionItem)

type CompletionContext struct {
	BaseContext

	// File *parser.File
	Position ts.Point

	// Publish a completion item
	Publish CompletionPublish
}

// Interface that supports completion can implement
type CompletionProvider interface {
	// Resolve go to definition
	ResolveCompletion(CompletionContext)
}
