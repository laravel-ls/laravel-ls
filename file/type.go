package file

import "strings"

type Type int

const (
	TypeUnknown = Type(iota)
	TypePHP
	TypeBlade
	TypeEnv
)

// Find the filetype based on filename
func TypeByFilename(filename string) Type {
	if strings.HasSuffix(filename, ".blade.php") {
		return TypeBlade
	}
	if strings.HasSuffix(filename, ".php") {
		return TypePHP
	}
	if strings.HasSuffix(filename, ".env") {
		return TypeEnv
	}
	return TypeUnknown
}
