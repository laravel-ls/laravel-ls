package env

import (
	"strings"

	"laravel-ls/parser"
	"laravel-ls/parser/env"
)

type Repository struct {
	data map[string]env.Metadata
}

func (r *Repository) Load(file *parser.File) error {
	r.Clear()
	data, err := env.Parse(file)
	if err == nil {
		r.data = data
	}
	return err
}

func (r Repository) Find(key string) map[string]env.Metadata {
	res := map[string]env.Metadata{}
	for k, v := range r.data {
		if strings.HasPrefix(k, key) {
			res[k] = v
		}
	}
	return res
}

func (r Repository) Get(key string) (meta env.Metadata, found bool) {
	meta, found = r.data[key]
	return
}

func (r Repository) Exists(key string) bool {
	_, found := r.Get(key)
	return found
}

func (r *Repository) Clear() {
	r.data = map[string]env.Metadata{}
}
