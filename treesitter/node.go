package treesitter

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

func FirstNamedChild(node *ts.Node) *ts.Node {
	if node.NamedChildCount() > 0 {
		return node.NamedChild(0)
	}
	return nil
}

func FirstNamedChildOfKind(node *ts.Node, kind string) *ts.Node {
	node = FirstNamedChild(node)
	if node != nil && node.Kind() == kind {
		return node
	}
	return nil
}

func NamedChildOfKind(node *ts.Node, kind string) *ts.Node {
	for i := uint(0); i < node.NamedChildCount(); i++ {
		child := node.NamedChild(i)
		if child.Kind() == kind {
			return child
		}
	}
	return nil
}
