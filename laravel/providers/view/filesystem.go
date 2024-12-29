package view

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"laravel-ls/file"
	"laravel-ls/laravel"
)

// TODO: Cache
type filesystem struct {
	// Directories where views are located
	paths []string
}

func (filesystem *filesystem) findInDir(basePath, dir string) ([]laravel.ViewFile, error) {
	files := []laravel.ViewFile{}

	root := path.Join(basePath, dir)

	err := filepath.Walk(root, func(filename string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && file.TypeByFilename(filename) == file.TypeBlade {
			filename, _ = filepath.Rel(root, filename)
			files = append(files, laravel.ViewFromPath(dir, filename))
		}
		return nil
	})

	return files, err
}

// Returns the directory where the file was found.
func (fs *filesystem) findDir(basePath, filename string) (string, bool) {
	for _, currentPath := range fs.paths {
		fs, err := os.Stat(path.Join(basePath, currentPath, filename))
		if err == nil && !fs.IsDir() {
			return currentPath, true
		}
	}
	return "", false
}

// Find a file in any of the directories.
func (fs *filesystem) find(basePath, filename string) (string, bool) {
	for _, currentPath := range fs.paths {
		fullPath := path.Join(basePath, currentPath, filename)
		fs, err := os.Stat(fullPath)
		if err == nil && !fs.IsDir() {
			return path.Join(currentPath, filename), true
		}
	}
	return "", false
}

func (fs *filesystem) search(basePath, search string) ([]laravel.ViewFile, error) {
	files := []laravel.ViewFile{}
	for _, currentPath := range fs.paths {
		filesInDir, err := fs.findInDir(basePath, currentPath)
		if err != nil {
			return []laravel.ViewFile{}, err
		}
		files = append(files, filesInDir...)
	}

	results := []laravel.ViewFile{}
	for _, file := range files {
		if len(search) < 1 || strings.HasPrefix(file.Name(), search) {
			results = append(results, file)
		}
	}
	return results, nil
}

// Find a view file on the filesystem
func (v *filesystem) findView(basePath, name string) (laravel.ViewFile, bool) {
	viewFile := laravel.ViewFromName(name)
	if dir, ok := v.findDir(basePath, viewFile.Filename()); ok {
		return laravel.ViewFromPath(dir, viewFile.Filename()), true
	}
	return laravel.ViewFile{}, false
}

func (fs *filesystem) exists(basepath, name string) bool {
	_, found := fs.findView(basepath, name)
	return found
}
