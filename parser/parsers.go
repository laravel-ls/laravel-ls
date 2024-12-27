package parser

import (
	"github.com/shufflingpixels/laravel-ls/file"
	"github.com/shufflingpixels/laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

var type_to_lang map[file.Type]*ts.Language = map[file.Type]*ts.Language{
	file.TypeUnknown: nil,
	file.TypePHP:     treesitter.GetLanguage(treesitter.LanguagePhp),
	file.TypeBlade:   treesitter.GetLanguage(treesitter.LanguageBlade),
	file.TypeEnv:     treesitter.GetLanguage(treesitter.LanguageDotEnv),
}

var parsers map[file.Type]*ts.Parser = map[file.Type]*ts.Parser{
	file.TypeUnknown: nil,
	file.TypePHP:     newParser(type_to_lang[file.TypePHP]),
	file.TypeBlade:   newParser(type_to_lang[file.TypeBlade]),
	file.TypeEnv:     newParser(type_to_lang[file.TypeEnv]),
}

func newParser(lang *ts.Language) *ts.Parser {
	parser := ts.NewParser()
	parser.SetLanguage(lang)
	return parser
}
