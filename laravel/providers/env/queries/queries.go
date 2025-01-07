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

func EnvCalls(file *parser.File) treesitter.CaptureSlice {
	a, _ := file.FindCaptures(treesitter.LanguagePhp, calls_query, QUERY_CAPTURE_ENV_KEY)
	b, _ := file.FindCaptures(treesitter.LanguagePhpOnly, calls_query, QUERY_CAPTURE_ENV_KEY)
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
