package env_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/parser/env"

	"github.com/stretchr/testify/require"
)

func TestParser_Basic(t *testing.T) {
	src := []byte(`
MAIL_MAILER=smtp
MAIL_HOST=mailpit
MAIL_PORT=1025
#comment
MAIL_USERNAME=null #comment
MAIL_PASSWORD=null
MAIL_ENCRYPTION=null
MAIL_FROM_ADDRESS="hello@example.com"
MAIL_FROM_NAME="${APP_NAME}"`)

	expected := map[string]env.Metadata{
		"MAIL_MAILER":       {Value: "smtp", Line: 1, Column: 12},
		"MAIL_HOST":         {Value: "mailpit", Line: 2, Column: 10},
		"MAIL_PORT":         {Value: "1025", Line: 3, Column: 10},
		"MAIL_USERNAME":     {Value: "null", Line: 5, Column: 14},
		"MAIL_PASSWORD":     {Value: "null", Line: 6, Column: 14},
		"MAIL_ENCRYPTION":   {Value: "null", Line: 7, Column: 16},
		"MAIL_FROM_ADDRESS": {Value: "hello@example.com", Line: 8, Column: 18},
		"MAIL_FROM_NAME":    {Value: "${APP_NAME}", Line: 9, Column: 15},
	}

	pFile, err := parser.Parse(src, file.TypeEnv)
	require.NoError(t, err)
	actual, err := env.Parse(pFile)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestParser_VariableSubstitution(t *testing.T) {
	src := []byte(`
APP_NAME=some_name

# Comment
MAIL_MAILER=smtp
MAIL_HOST=mailpit
MAIL_PORT=1025
MAIL_USERNAME=null
MAIL_PASSWORD=null
MAIL_ENCRYPTION="encrypt:${APP_NAME}" #comment
MAIL_FROM_ADDRESS="hello@example.com"
MAIL_FROM_NAME="xxx ${MAIL_ENCRYPTION} yyy"`)

	expected := map[string]env.Metadata{
		"APP_NAME":          {Value: "some_name", Line: 1, Column: 9},
		"MAIL_MAILER":       {Value: "smtp", Line: 4, Column: 12},
		"MAIL_HOST":         {Value: "mailpit", Line: 5, Column: 10},
		"MAIL_PORT":         {Value: "1025", Line: 6, Column: 10},
		"MAIL_USERNAME":     {Value: "null", Line: 7, Column: 14},
		"MAIL_PASSWORD":     {Value: "null", Line: 8, Column: 14},
		"MAIL_ENCRYPTION":   {Value: "encrypt:some_name", Line: 9, Column: 16},
		"MAIL_FROM_ADDRESS": {Value: "hello@example.com", Line: 10, Column: 18},
		"MAIL_FROM_NAME":    {Value: "xxx encrypt:some_name yyy", Line: 11, Column: 15},
	}

	pFile, err := parser.Parse(src, file.TypeEnv)
	require.NoError(t, err)
	actual, err := env.Parse(pFile)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestParser_VariableSubstitutionDefinedAfter(t *testing.T) {
	src := []byte(`
MAIL_FROM_NAME="xxx ${APP_NAME} yyy"
APP_NAME=some_name`)

	expected := map[string]env.Metadata{
		"MAIL_FROM_NAME": {Value: "xxx ${APP_NAME} yyy", Line: 1, Column: 15},
		"APP_NAME":       {Value: "some_name", Line: 2, Column: 9},
	}

	pFile, err := parser.Parse(src, file.TypeEnv)
	require.NoError(t, err)
	actual, err := env.Parse(pFile)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
