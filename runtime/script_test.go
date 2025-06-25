package runtime_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/laravel-ls/laravel-ls/runtime"
	"github.com/stretchr/testify/require"
)

type MockProcess struct {
	response []byte
}

func (p MockProcess) Exec(workingDir string, code []byte) (io.Reader, error) {
	return bytes.NewReader(p.response), nil
}

func Test_JsonError(t *testing.T) {
	proc := MockProcess{
		response: []byte("Non json"),
	}

	type result struct{}

	r := result{}

	ret, err := runtime.CallScript(proc, "/tmp/root", []byte("code"), &r)
	require.EqualError(t, err, "json: invalid character 'N' looking for beginning of value")
	require.Equal(t, ret, &r)
}

func Test_CorrectJson(t *testing.T) {
	proc := MockProcess{
		response: []byte(`[{"value":"value1"}, {"value":"value2"}]`),
	}

	type result struct {
		Value string `json:"value"`
	}

	r := []result{}
	expected := []result{
		{Value: "value1"},
		{Value: "value2"},
	}

	ret, err := runtime.CallScript(proc, "/tmp/root", []byte("code"), &r)
	require.NoError(t, err)
	require.Equal(t, &r, ret)
	require.Equal(t, ret, &expected)
}
