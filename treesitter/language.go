package treesitter

import (
	blade "github.com/EmranMR/tree-sitter-blade/bindings/go"
	dotenv "github.com/pnx/tree-sitter-dotenv/bindings/go"
	"laravel-ls/file"
	ts "github.com/tree-sitter/go-tree-sitter"
	html "github.com/tree-sitter/tree-sitter-html/bindings/go"
	php "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

// Constants for all treesitter languages that are supported
const (
	LangaugeUnknown = "unknown"
	LanguageHTML    = "html"
	LanguagePhp     = "php"
	LanguagePhpOnly = "php_only"
	LanguageBlade   = "blade"
	LanguageDotEnv  = "dotenv"
)

// Holds a map of language strings and treesitter language objects.
var langmap map[string]*ts.Language = map[string]*ts.Language{
	LangaugeUnknown: nil,
	LanguageHTML:    ts.NewLanguage(html.Language()),
	LanguagePhp:     ts.NewLanguage(php.LanguagePHP()),
	LanguagePhpOnly: ts.NewLanguage(php.LanguagePHPOnly()),
	LanguageBlade:   ts.NewLanguage(blade.Language()),
	LanguageDotEnv:  ts.NewLanguage(dotenv.Language()),
}

func GetLanguage(lang string) *ts.Language {
	if lang, ok := langmap[lang]; ok {
		return lang
	}
	return nil
}

// Get the language for a particular filetype
func FiletypeToLanguage(t file.Type) string {
	switch t {
	case file.TypePHP:
		return LanguagePhp
	case file.TypeBlade:
		return LanguageBlade
	case file.TypeEnv:
		return LanguageDotEnv
	}
	return LangaugeUnknown
}
