package assets

import (
	"embed"
	"path"
)

//go:embed queries/*/*.scm
var FS embed.FS

func QueryPath(filename string) string {
	return path.Join("queries", filename)
}
