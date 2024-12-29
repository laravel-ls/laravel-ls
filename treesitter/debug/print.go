package debug

import (
	"strings"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type Print struct {
	AnonNodes  bool
	IndentSize uint
}

func (p Print) Indent(depth uint) string {
	return strings.Repeat(" ", int(p.IndentSize*depth))
}

func (p Print) Inner(node *ts.Node, depth uint) string {
	ret := p.Indent(depth) + FormatNode(node) + "\n"

	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		if p.AnonNodes == false && child.IsNamed() == false {
			continue
		}
		ret += p.Inner(node, depth+1)
	}

	return ret
}

func (p Print) Print(node *ts.Node) string {
	return p.Inner(node, 0)
}
