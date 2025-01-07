package parser

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Capture struct {
	Name string
	Node ts.Node
}
