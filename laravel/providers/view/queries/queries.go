package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/queries"

	ts "github.com/tree-sitter/go-tree-sitter"
)

const QUERY_CAPTURE_VIEW_NAME = "view.name"

func getQuery() string {
	q, _ := queries.GetQuery(treesitter.LanguagePhp, "view")
	return q
}

// Check if node is a view name.
func IsViewName(file *parser.File, node *ts.Node) bool {
	return file.NodeMatchesCapture(treesitter.LanguagePhp, getQuery(), QUERY_CAPTURE_VIEW_NAME, node)
}

func ViewNames(file *parser.File) treesitter.CaptureSlice {
	query := getQuery()
	a, _ := file.FindCaptures(treesitter.LanguagePhp, query, QUERY_CAPTURE_VIEW_NAME)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, query, QUERY_CAPTURE_VIEW_NAME)
	return append(a, b...)
}
