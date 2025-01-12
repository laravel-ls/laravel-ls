package queries

import (
	"embed"
	"errors"
	"io/fs"
	"path"
)

//go:embed */*.scm
var query embed.FS

var ErrQueryNotFound error = errors.New("query not found")

func GetQuery(lang, name string) (string, error) {
	q, err := query.ReadFile(path.Join(lang, name+".scm"))

	// if file does not exist, return ErrQueryNotFound.
	if err != nil && errors.Unwrap(err) == fs.ErrNotExist {
		err = ErrQueryNotFound
	}
	return string(q), err
}

func GetInjectionQuery(lang string) (string, error) {
	return GetQuery(lang, "injections")
}
