package queries

import (
	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
)

const QueryCaptureAppFilename = "app.service"

func queryAppCalls(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	r, err := file.FindTags(lang, QueryCaptureAppFilename)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func AppCalls(file *parser.File) treesitter.CaptureSlice {
	return append(queryAppCalls(file, language.PHP),
		queryAppCalls(file, language.PHPOnly)...)
}
