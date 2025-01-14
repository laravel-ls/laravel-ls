package treesitter

import (
	"errors"
	"io/fs"
	"path"

	"github.com/laravel-ls/laravel-ls/treesitter/assets"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// A map of "virtual files" that points to some file that actually exist.
// this is a hack that's needed because php and php_only are different
// languages in treesitter (but are logically the same).
var fileAlias = map[string]string{
	"php_only/env.scm":  "php/env.scm",
	"php_only/view.scm": "php/view.scm",
}

var ErrQueryNotFound error = errors.New("query not found")

// Read a query from file
func ReadQueryFromFile(lang, name string) (string, error) {
	filename := path.Join(lang, name+".scm")

	// Resolve any alias
	if alias, ok := fileAlias[filename]; ok {
		filename = alias
	}

	source, err := assets.FS.ReadFile(assets.QueryPath(filename))

	// if file does not exist, return ErrQueryNotFound.
	if err != nil && errors.Unwrap(err) == fs.ErrNotExist {
		err = ErrQueryNotFound
	}
	return string(source), err
}

func GetQuery(lang, name string) (*ts.Query, error) {
	tsLang := GetLanguage(lang)
	if tsLang == nil {
		return nil, ErrLangNotSupported
	}

	source, err := ReadQueryFromFile(lang, name)
	if err != nil {
		return nil, err
	}

	query, tsErr := ts.NewQuery(tsLang, source)
	// This error checking is needed because of:
	// https://go.dev/doc/faq#nil_error
	if tsErr != nil {
		return nil, tsErr
	}
	return query, nil
}

func GetInjectionQuery(lang string) (*ts.Query, error) {
	return GetQuery(lang, "injections")
}
