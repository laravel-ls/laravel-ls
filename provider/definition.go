package provider

import (
	"github.com/laravel-ls/protocol"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type LocationPublisher func(protocol.Location)

type DefinitionContext struct {
	BaseContext

	Position ts.Point

	// Publish a location
	Publish LocationPublisher
}

// Interface that providers that supports go to definition can implement
type DefinitionProvider interface {
	// Resolve go to definition
	ResolveDefinition(DefinitionContext)
}
