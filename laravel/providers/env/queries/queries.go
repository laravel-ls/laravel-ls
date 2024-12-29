package queries

import (
	_ "embed"

	"github.com/shufflingpixels/laravel-ls/parser"
	"github.com/shufflingpixels/laravel-ls/treesitter"
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

func GetKey(node *ts.Node, source []byte) string {
	if (node.Kind() == "string" || node.Kind() == "encapsed_string") && node.NamedChildCount() > 0 {
		contentNode := node.NamedChild(0)
		if contentNode.Kind() == "string_content" {
			return contentNode.Utf8Text(source)
		}
	}
	return ""
}
