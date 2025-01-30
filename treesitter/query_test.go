package treesitter_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
	"github.com/stretchr/testify/assert"
)

func TestGetInjectionQuery(t *testing.T) {
	_, err := treesitter.GetInjectionQuery(language.PHP)
	assert.NoError(t, err)

	_, err = treesitter.GetInjectionQuery(language.PHPOnly)
	assert.Error(t, err)
}
