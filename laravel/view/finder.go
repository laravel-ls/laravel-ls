package view

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// Find view files on the filesystem.
// This struct is based on the Illuminate\View\FileViewFinder
// class from laravel source.
type Finder struct {
	// Directories where views are located
	paths []string

	// List of file extensions views can have.
	extensions []string

	// Filesystem
	fs afero.Fs
}

func NewFinder(fs afero.Fs) *Finder {
	return &Finder{
		fs: fs,
		extensions: []string{
			".blade.php",
			".php",
		},
	}
}

func (finder *Finder) AddLocation(path string) {
	finder.paths = append(finder.paths, path)
}

func (finder Finder) Paths() []string {
	return finder.paths
}

func (finder *Finder) RegisterExtensions(extensions ...string) {
	finder.extensions = extensions
}

func (finder Finder) Extensions() []string {
	return finder.extensions
}

// Get the fully qualified location of the view.
// Returns false if no location is found.
func (finder Finder) Find(name string) (string, bool) {
	for _, path := range finder.PossibleFiles(name) {
		if stat, err := finder.fs.Stat(path); err == nil && !stat.IsDir() {
			return path, true
		}
	}
	return "", false
}

// Returns all possible files for a given view name.
func (finder Finder) PossibleFiles(name string) []string {
	filenames := finder.getAllFilenames(name)
	files := []string{}
	for _, base := range finder.paths {
		for _, filename := range filenames {
			files = append(files, filepath.Join(base, filename))
		}
	}
	return files
}

// Lists all view files containing the input string in their name.
func (finder Finder) Search(input string) []View {
	var matches []View

	for _, basePath := range finder.paths {
		_ = afero.Walk(finder.fs, basePath, func(fullPath string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}

			if ext := finder.matchingExtension(fullPath); ext != "" {
				name := finder.viewNameFromPath(basePath, fullPath, ext)
				if strings.Contains(name, input) {
					matches = append(matches, View{
						name: name,
						path: fullPath,
					})
				}
			}
			return nil
		})
	}

	return matches
}

func (finder Finder) matchingExtension(path string) string {
	for _, ext := range finder.extensions {
		if strings.HasSuffix(path, ext) {
			return ext
		}
	}
	return ""
}

func (finder Finder) viewNameFromPath(basePath, fullPath, ext string) string {
	rel, err := filepath.Rel(basePath, fullPath)
	if err != nil {
		return ""
	}
	name := strings.TrimSuffix(rel, ext)
	return strings.ReplaceAll(name, string(os.PathSeparator), ".")
}

func (finder Finder) getAllFilenames(name string) []string {
	name = strings.ReplaceAll(name, ".", string(os.PathSeparator))
	files := []string{}
	for _, ext := range finder.extensions {
		files = append(files, name+ext)
	}
	return files
}
