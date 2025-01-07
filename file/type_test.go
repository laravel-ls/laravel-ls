package file_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/file"
)

func Test_TypeByFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected file.Type
	}{
		{
			name:     "Blade",
			filename: "/path/to/file.blade.php",
			expected: file.TypeBlade,
		},
		{
			name:     "PHP",
			filename: "/path/to/file.php",
			expected: file.TypePHP,
		},
		{
			name:     "Env",
			filename: "/path/to/.env",
			expected: file.TypeEnv,
		},
		{
			name:     "Env Example",
			filename: "/path/to/.env.example",
			expected: file.TypeEnv,
		},
		{
			name:     "Env local",
			filename: "/path/to/.env.local",
			expected: file.TypeEnv,
		},
	}

	for _, args := range tests {
		t.Run(args.name, func(t *testing.T) {
			actual := file.TypeByFilename(args.filename)
			if actual != args.expected {
				t.Errorf("%v is not equal to %v", actual, args.expected)
			}
		})
	}
}
