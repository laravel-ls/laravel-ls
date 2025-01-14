package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
)

const QUERY_CAPTURE_VIEW_NAME = "view.name"

func getQuery() string {
	q, _ := treesitter.GetQuery(treesitter.LanguagePhp, "view")
	return q
}

func ViewNames(file *parser.File) treesitter.CaptureSlice {
	query := getQuery()
	a, _ := file.FindCaptures(treesitter.LanguagePhp, query, QUERY_CAPTURE_VIEW_NAME)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, query, QUERY_CAPTURE_VIEW_NAME)
	return append(a, b...)
}
