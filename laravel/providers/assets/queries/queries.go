package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
)

const QueryCaptureAssetFilename = "asset.filename"

func queryAssetCalls(file *parser.File, lang string) treesitter.CaptureSlice {
	query, err := treesitter.GetQuery(lang, "asset")
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	r, err := file.FindCaptures(lang, query, QueryCaptureAssetFilename)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func Assets(file *parser.File) treesitter.CaptureSlice {
	return append(queryAssetCalls(file, treesitter.LanguagePhp),
		queryAssetCalls(file, treesitter.LanguagePhpOnly)...)
}
