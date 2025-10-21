package file

import (
	"path"
	"strings"
)

type Type int

const (
	TypeUnknown = Type(iota)
	TypePHP
	TypeBlade
	TypeEnv
)

type checkFn func(string) bool

func hasExtension(ext string) checkFn {
	return func(filename string) bool {
		return strings.HasSuffix(filename, ext)
	}
}

func envFile(filename string) bool {
	return filename == ".env" || strings.HasPrefix(filename, ".env.")
}

// TypeByFilename finds the filetype based on filename
func TypeByFilename(filename string) Type {
	lookup := map[Type]checkFn{
		TypeBlade: hasExtension(".blade.php"),
		TypePHP:   hasExtension(".php"),
		TypeEnv:   envFile,
	}

	filename = path.Base(filename)
	for typ, test := range lookup {
		if ok := test(filename); ok {
			return typ
		}
	}
	return TypeUnknown
}
