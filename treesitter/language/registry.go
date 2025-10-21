package language

import (
	"errors"

	blade "github.com/EmranMR/tree-sitter-blade/bindings/go"
	"github.com/laravel-ls/laravel-ls/file"
	dotenv "github.com/pnx/tree-sitter-dotenv/bindings/go"
	html "github.com/tree-sitter/tree-sitter-html/bindings/go"
	php "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

var ErrNotSupported = errors.New("language is not supported")

const (
	HTML    Identifier = "html"
	PHP     Identifier = "php"
	PHPOnly Identifier = "php_only"
	Blade   Identifier = "blade"
	DotEnv  Identifier = "dotenv"
)

// Holds a map of identifiers and language objects.
var langmap map[Identifier]*Language = map[Identifier]*Language{
	HTML:    NewLanguage(HTML, html.Language()),
	PHP:     NewLanguage(PHP, php.LanguagePHP()),
	PHPOnly: NewLanguage(PHPOnly, php.LanguagePHPOnly()),
	Blade:   NewLanguage(Blade, blade.Language()),
	DotEnv:  NewLanguage(DotEnv, dotenv.Language()),
}

func Get(lang Identifier) *Language {
	if lang, ok := langmap[lang]; ok {
		return lang
	}
	return nil
}

// FiletypeToLanguage Get the language Identifier for a particular filetype
func FiletypeToLanguage(t file.Type) Identifier {
	switch t {
	case file.TypePHP:
		return PHP
	case file.TypeBlade:
		return Blade
	case file.TypeEnv:
		return DotEnv
	}
	return InvalidIdentifier
}
