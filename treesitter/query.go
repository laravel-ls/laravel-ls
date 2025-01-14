package treesitter

import (
	"errors"
	"io/fs"
	"path"

	"github.com/laravel-ls/laravel-ls/treesitter/assets"
)

var ErrQueryNotFound error = errors.New("query not found")

func GetQuery(lang, name string) (string, error) {
	q, err := assets.FS.ReadFile(assets.QueryPath(path.Join(lang, name+".scm")))

	// if file does not exist, return ErrQueryNotFound.
	if err != nil && errors.Unwrap(err) == fs.ErrNotExist {
		err = ErrQueryNotFound
	}
	return string(q), err
}

func GetInjectionQuery(lang string) (string, error) {
	return GetQuery(lang, "injections")
}
