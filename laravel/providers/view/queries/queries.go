package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

const QUERY_CAPTURE_VIEW_NAME = "view.name"

//go:embed view.scm
var view_query string

// Check if node is a view name.
func IsViewName(file *parser.File, node *ts.Node) bool {
	return file.NodeMatchesCapture(treesitter.LanguagePhp, view_query, QUERY_CAPTURE_VIEW_NAME, node)
}

func ViewNameAtPosition(file *parser.File, position ts.Point) *ts.Node {
	for _, capture := range ViewNames(file) {
		if treesitter.PointInRange(position, capture.Node.Range()) {
			return &capture.Node
		}
	}
	return nil
}

func ViewNames(file *parser.File) []parser.Capture {
	a, _ := file.FindCaptures(treesitter.LanguagePhp, view_query, QUERY_CAPTURE_VIEW_NAME)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, view_query, QUERY_CAPTURE_VIEW_NAME)
	return append(a, b...)
}

func GetViewName(node *ts.Node, source []byte) string {
	if (node.Kind() == "string" || node.Kind() == "encapsed_string") && node.NamedChildCount() > 0 {
		contentNode := node.NamedChild(0)
		if contentNode.Kind() == "string_content" {
			return contentNode.Utf8Text(source)
		}
	}
	return ""
}
