package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
)

const QueryCaptureViewName = "view.name"

func findViewNames(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	query, err := treesitter.GetQuery(lang, "view")
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	r, err := file.FindCaptures(lang, query, QueryCaptureViewName)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func ViewNames(file *parser.File) treesitter.CaptureSlice {
	return append(findViewNames(file, language.PHP),
		findViewNames(file, language.PHPOnly)...)
}
