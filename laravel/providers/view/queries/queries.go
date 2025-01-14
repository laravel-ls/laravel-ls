package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
)

const QueryCaptureViewName = "view.name"

func findViewNames(file *parser.File, lang string) treesitter.CaptureSlice {
	query, err := treesitter.GetQuery(lang, "view")
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	defer query.Close()
	r, err := file.FindCaptures(lang, query, QueryCaptureViewName)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func ViewNames(file *parser.File) treesitter.CaptureSlice {
	return append(findViewNames(file, treesitter.LanguagePhp),
		findViewNames(file, treesitter.LanguagePhpOnly)...)
}
