package env

import (
	"strings"

	"github.com/laravel-ls/laravel-ls/env"
	"github.com/laravel-ls/laravel-ls/env/evaluator"
	"github.com/laravel-ls/laravel-ls/parser"
)

type Repository struct {
	variables map[string]env.Variable
}

func (r *Repository) Load(file *parser.File) error {
	r.Clear()
	data, err := evaluator.Evaluate(file)
	if err == nil {
		r.variables = data
	}
	return err
}

func (r Repository) Find(key string) map[string]env.Variable {
	res := map[string]env.Variable{}
	for k, v := range r.variables {
		if strings.HasPrefix(k, key) {
			res[k] = v
		}
	}
	return res
}

func (r Repository) Get(key string) (env.Variable, bool) {
	variable, found := r.variables[key]
	return variable, found
}

func (r Repository) Exists(key string) bool {
	_, found := r.Get(key)
	return found
}

func (r *Repository) Clear() {
	r.variables = map[string]env.Variable{}
}
