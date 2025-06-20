package asset

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/afero"
)

type Finder struct {
	rootPath string
	fs       afero.Fs
}

func NewFinder(fs afero.Fs, rootPath string) *Finder {
	return &Finder{
		rootPath: path.Join(rootPath, "public"),
		fs:       fs,
	}
}

func (finder Finder) Search(search string) []string {
	files := []string{}

	afero.Walk(finder.fs, finder.rootPath, func(fullPath string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.Contains(fullPath, search) {
			files = append(files, fullPath)
		}
		return nil
	})

	return files
}

func (finder Finder) Exists(filename string) bool {
	stat, err := finder.fs.Stat(path.Join(finder.rootPath, filename))
	return err == nil && !stat.IsDir()
}
