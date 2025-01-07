package assets

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/laravel-ls/laravel-ls/file"
)

type Repository struct{}

func (r Repository) findInDir(basePath, dir string) ([]string, error) {
	files := []string{}

	root := path.Join(basePath, dir)

	err := filepath.Walk(root, func(filename string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && file.TypeByFilename(filename) != file.TypePHP {
			filename, _ = filepath.Rel(root, filename)
			files = append(files, filename)
		}
		return nil
	})

	return files, err
}

func (r Repository) Search(basePath, search string) []string {
	files, err := r.findInDir(basePath, "public")
	if err != nil {
		return []string{}
	}

	results := []string{}
	for _, filename := range files {
		if len(search) < 1 || strings.HasPrefix(filename, search) {
			results = append(results, filename)
		}
	}
	return results
}

func (r Repository) Exists(basePath, filename string) bool {
	info, err := os.Stat(path.Join(basePath, "public", filename))
	return err == nil && !info.IsDir()
}
