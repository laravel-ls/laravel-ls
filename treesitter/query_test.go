package treesitter_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/stretchr/testify/assert"
)

func TestGetInjectionQuery(t *testing.T) {
	_, err := treesitter.GetInjectionQuery("php")
	assert.NoError(t, err)

	// It be weird if this language is included.
	_, err = treesitter.GetInjectionQuery("pascal")
	assert.Error(t, err)
}
