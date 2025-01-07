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

func ViewNames(file *parser.File) parser.CaptureSlice {
	a, _ := file.FindCaptures(treesitter.LanguagePhp, view_query, QUERY_CAPTURE_VIEW_NAME)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, view_query, QUERY_CAPTURE_VIEW_NAME)
	return append(a, b...)
}
