package provider

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Hover struct {
	Content string
	Range   *ts.Range
}

type HoverPublisher func(Hover)

type HoverContext struct {
	BaseContext
	Position ts.Point
	Publish  HoverPublisher
}

// Interface that plugins that supports hover information can implement
type HoverProvider interface {
	Hover(HoverContext)
}
