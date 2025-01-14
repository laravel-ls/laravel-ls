package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

const QUERY_CAPTURE_ENV_KEY = "env.key"

func queryEnvCalls(file *parser.File, lang string) treesitter.CaptureSlice {
	query, err := treesitter.GetQuery(lang, "env")
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	defer query.Close()
	r, err := file.FindCaptures(lang, query, QUERY_CAPTURE_ENV_KEY)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func EnvCalls(file *parser.File) treesitter.CaptureSlice {
	return append(queryEnvCalls(file, treesitter.LanguagePhp),
		queryEnvCalls(file, treesitter.LanguagePhpOnly)...)
}

// Check if a env call has a default value.
func HasDefault(node *ts.Node) bool {
	parent := node.Parent()

	if parent != nil && parent.Kind() == "argument" {
		return parent.NextNamedSibling() != nil
	}
	return false
}
