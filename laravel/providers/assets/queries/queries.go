package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
)

const QueryCaptureAssetFilename = "asset.filename"

func queryAssetCalls(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	r, err := file.FindTags(lang, QueryCaptureAssetFilename)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func Assets(file *parser.File) treesitter.CaptureSlice {
	return append(queryAssetCalls(file, language.PHP),
		queryAssetCalls(file, language.PHPOnly)...)
}
