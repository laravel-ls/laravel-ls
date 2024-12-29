package transport_test

import (
	"os"
	"testing"

	"laravel-ls/lsp/transport"
	"github.com/stretchr/testify/assert"
)

func Test_Read(t *testing.T) {
	r, w, err := os.Pipe()
	defer r.Close()
	defer w.Close()
	assert.NoError(t, err)

	os.Stdin = r

	stdio := transport.NewStdio()

	b, err := w.Write([]byte("Hello World\n"))
	assert.Equal(t, 12, b)
	assert.NoError(t, err)

	actual := make([]byte, 12)
	b, err = stdio.Read(actual)
	assert.NoError(t, err)
	assert.Equal(t, 12, b)
	assert.Equal(t, []byte("Hello World\n"), actual)
}

func Test_Write(t *testing.T) {
	r, w, err := os.Pipe()
	defer r.Close()
	defer w.Close()
	assert.NoError(t, err)

	os.Stdout = w

	stdio := transport.NewStdio()

	b, err := stdio.Write([]byte("Hello World\n"))
	assert.Equal(t, 12, b)
	assert.NoError(t, err)

	actual := make([]byte, 12)
	b, err = r.Read(actual)
	assert.NoError(t, err)
	assert.Equal(t, 12, b)
	assert.Equal(t, []byte("Hello World\n"), actual)
}
