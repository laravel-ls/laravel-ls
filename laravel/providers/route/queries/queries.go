package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
)

const QueryCaptureRouteKey = "route.name"

func queryRouteCalls(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	r, err := file.FindTags(lang, QueryCaptureRouteKey)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func RouteCalls(file *parser.File) treesitter.CaptureSlice {
	return append(queryRouteCalls(file, language.PHP),
		queryRouteCalls(file, language.PHPOnly)...)
}
