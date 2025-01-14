package cache

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/parser"
)

// FileCache holds information about open files in the workspace.
type FileCache struct {
	// Map of files in the cache.
	files map[string]*parser.File
}

func NewFileCache() *FileCache {
	return &FileCache{
		files: map[string]*parser.File{},
	}
}

func (s *FileCache) Open(filename string) (*parser.File, error) {
	// If file is already opened, return it.
	if file := s.Get(filename); file != nil {
		return file, nil
	}

	lang := file.TypeByFilename(filename)
	if lang == file.TypeUnknown {
		return nil, errors.New("language not supported")
	}

	content, err := s.read(filename)
	if err != nil {
		return nil, err
	}

	parseFile, err := parser.Parse(content, lang)
	if err == nil {
		s.files[filename] = parseFile
	}
	return parseFile, err
}

func (s *FileCache) read(filename string) ([]byte, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	defer fd.Close()

	return io.ReadAll(fd)
}

func (s FileCache) Get(filename string) *parser.File {
	if file, ok := s.files[filename]; ok {
		return file
	}
	return nil
}

func (s *FileCache) Close(filename string) error {
	if s.IsOpen(filename) {
		delete(s.files, filename)
		return nil
	}
	return fmt.Errorf("file '%s' is not open", filename)
}

func (s FileCache) IsOpen(filename string) bool {
	_, ok := s.files[filename]
	return ok
}
