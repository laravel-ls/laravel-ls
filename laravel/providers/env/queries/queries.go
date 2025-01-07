package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

const QUERY_CAPTURE_ENV_KEY = "env.key"

//go:embed calls.scm
var calls_query string

func EnvCalls(file *parser.File) []parser.Capture {
	a, _ := file.FindCaptures(treesitter.LanguagePhp, calls_query, QUERY_CAPTURE_ENV_KEY)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, calls_query, QUERY_CAPTURE_ENV_KEY)
	return append(a, b...)
}

func EnvCallAtPosition(file *parser.File, position ts.Point) *ts.Node {
	for _, capture := range EnvCalls(file) {
		if treesitter.PointInRange(position, capture.Node.Range()) {
			return &capture.Node
		}
	}
	return nil
}

func EnvCallsInRange(file *parser.File, r ts.Range) []*ts.Node {
	nodes := []*ts.Node{}
	for _, capture := range EnvCalls(file) {
		if treesitter.RangeOverlap(r, capture.Node.Range()) {
			nodes = append(nodes, &capture.Node)
		}
	}
	return nodes
}

func GetKey(node *ts.Node, source []byte) string {
	if (node.Kind() == "string" || node.Kind() == "encapsed_string") && node.NamedChildCount() > 0 {
		contentNode := node.NamedChild(0)
		if contentNode.Kind() == "string_content" {
			return contentNode.Utf8Text(source)
		}
	}
	return ""
}

// Check if a env call has a default value.
func HasDefault(node *ts.Node) bool {
	parent := node.Parent()

	if parent != nil && parent.Kind() == "argument" {
		return parent.NextNamedSibling() != nil
	}
	return false
}
