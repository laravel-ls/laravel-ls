package treesitter

import (
	"errors"
	"fmt"
	"io/fs"
	"path"

	"github.com/laravel-ls/laravel-ls/treesitter/assets"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
	"github.com/laravel-ls/laravel-ls/utils/cache"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var queryCache = cache.New[*ts.Query]()

var ErrQueryNotFound error = errors.New("query not found")

// Read a query from file
func ReadQueryFromFile(lang *language.Language, name string) (string, error) {
	filename := path.Join(lang.Name(), name+".scm")

	// Hack to use the php tags file for php_only
	if filename == "php_only/tags.scm" {
		filename = "php/tags.scm"
	}

	source, err := assets.FS.ReadFile(assets.QueryPath(filename))

	// if file does not exist, return ErrQueryNotFound.
	if err != nil && errors.Unwrap(err) == fs.ErrNotExist {
		err = ErrQueryNotFound
	}
	return string(source), err
}

func GetQuery(lang_id language.Identifier, name string) (*ts.Query, error) {
	key := fmt.Sprintf("%s:%s", lang_id.String(), name)

	return queryCache.Remember(key, func(string) (*ts.Query, error) {
		lang := language.Get(lang_id)
		if lang == nil {
			return nil, language.ErrNotSupported
		}

		source, err := ReadQueryFromFile(lang, name)
		if err != nil {
			return nil, err
		}

		query, tsErr := lang.Query(source)
		// This error checking is needed because of:
		// https://go.dev/doc/faq#nil_error
		if tsErr != nil {
			return nil, tsErr
		}
		return query, nil
	})
}

func GetInjectionQuery(lang language.Identifier) (*ts.Query, error) {
	return GetQuery(lang, "injections")
}

func GetTagsQuery(lang language.Identifier) (*ts.Query, error) {
	return GetQuery(lang, "tags")
}

func FreeQueryCache() {
	for _, query := range queryCache.Items() {
		query.Close()
	}
	queryCache.Clear()
}
