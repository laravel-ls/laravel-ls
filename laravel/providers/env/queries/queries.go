package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

const QUERY_CAPTURE_ENV_KEY = "env.key"

func getQuery() string {
	q, _ := treesitter.GetQuery(treesitter.LanguagePhp, "env")
	return q
}

func EnvCalls(file *parser.File) treesitter.CaptureSlice {
	query := getQuery()
	a, _ := file.FindCaptures(treesitter.LanguagePhp, query, QUERY_CAPTURE_ENV_KEY)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, query, QUERY_CAPTURE_ENV_KEY)
	return append(a, b...)
}

// Check if a env call has a default value.
func HasDefault(node *ts.Node) bool {
	parent := node.Parent()

	if parent != nil && parent.Kind() == "argument" {
		return parent.NextNamedSibling() != nil
	}
	return false
}
