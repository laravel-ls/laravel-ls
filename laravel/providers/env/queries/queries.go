package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"

	ts "github.com/tree-sitter/go-tree-sitter"
)

const QueryCaptureEnvKey = "env.key"

func queryEnvCalls(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	r, err := file.FindTags(lang, QueryCaptureEnvKey)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func EnvCalls(file *parser.File) treesitter.CaptureSlice {
	return append(queryEnvCalls(file, language.PHP),
		queryEnvCalls(file, language.PHPOnly)...)
}

// HasDefault check if a env call has a default value.
func HasDefault(node *ts.Node) bool {
	parent := node.Parent()

	if parent != nil && parent.Kind() == "argument" {
		return parent.NextNamedSibling() != nil
	}
	return false
}
