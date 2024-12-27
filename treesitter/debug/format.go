package debug

import (
	"fmt"

	ts "github.com/tree-sitter/go-tree-sitter"
)

func FormatPoint(p ts.Point) string {
	return fmt.Sprintf("[%d %d]", p.Row, p.Column)
}

func FormatRange(r ts.Range) string {
	return fmt.Sprintf("%s - %s",
		FormatPoint(r.StartPoint),
		FormatPoint(r.EndPoint))
}

func FormatNode(node *ts.Node) string {
	return fmt.Sprintf("%s; %s", node.Kind(), FormatRange(node.Range()))
}
