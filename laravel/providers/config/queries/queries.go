package queries

import (
	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
)

const QueryCaptureConfigKey = "config.key"

func queryConfigCalls(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	r, err := file.FindTags(lang, QueryCaptureConfigKey)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func ConfigCalls(file *parser.File) treesitter.CaptureSlice {
	return append(queryConfigCalls(file, language.PHP),
		queryConfigCalls(file, language.PHPOnly)...)
}
