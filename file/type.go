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

// Find the filetype based on filename
func TypeByFilename(filename string) Type {
	lookup := map[string]Type{
		".blade.php": TypeBlade,
		".php":       TypePHP,
	}

	for ext, typ := range lookup {
		if strings.HasSuffix(filename, ext) {
			return typ
		}
	}

	if strings.HasPrefix(path.Base(filename), ".env") {
		return TypeEnv
	}

	return TypeUnknown
}
