package laravel

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// Struct to represent a view file.
type ViewFile struct {
	filename string

	// Directory relative to project root.
	// for example: resources/views or vendor/package/views
	directory string
}

func (v ViewFile) Name() string {
	name := strings.ReplaceAll(v.filename, string(os.PathSeparator), ".")
	name, _ = strings.CutSuffix(name, ".blade.php")
	return name
}

func (v ViewFile) Filename() string {
	return v.filename
}

func (v ViewFile) Directory() string {
	return v.directory
}

func (v ViewFile) Path() string {
	if len(v.Directory()) > 0 {
		return path.Join(v.Directory(), v.Filename())
	}
	return v.Filename()
}

func (v ViewFile) String() string {
	return v.Filename()
}

func ViewFromFilename(filename string) ViewFile {
	return ViewFile{
		filename: filename,
	}
}

func ViewFromPath(baseDir string, filename string) ViewFile {
	return ViewFile{
		directory: baseDir,
		filename:  filename,
	}
}

func ViewFromName(name string) ViewFile {
	filename := strings.ReplaceAll(name, ".", string(os.PathSeparator))
	filename = fmt.Sprintf("%s.blade.php", filename)
	return ViewFromFilename(filename)
}
