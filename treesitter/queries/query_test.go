package queries_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/treesitter/queries"
	"github.com/stretchr/testify/assert"
)

func TestGetInjectionQuery(t *testing.T) {
	_, err := queries.GetInjectionQuery("php")
	assert.NoError(t, err)

	// It be weird if this language is included.
	_, err = queries.GetInjectionQuery("pascal")
	assert.Error(t, err)
}
