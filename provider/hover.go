package provider

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Hover struct{}

type HoverContext struct {
	BaseContext

	Position ts.Point
}

// Interface that plugins that supports hover information can implement
type HoverProvider interface {
	Hover(HoverContext) string
}
