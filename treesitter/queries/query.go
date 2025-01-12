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

func GetInjectionQuery(lang string) (string, error) {
	q, err := query.ReadFile(path.Join(lang, "injections.scm"))

	// if file does not exist, return ErrQueryNotFound.
	if err != nil && errors.Unwrap(err) == fs.ErrNotExist {
		err = ErrQueryNotFound
	}
	return string(q), err
}
