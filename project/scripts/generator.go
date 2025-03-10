//go:build ignore

package main

import (
	"os"
	"strings"

	"github.com/laravel-ls/laravel-ls/runtime/template"
)

func isSourceFile(name string) bool {
	return strings.HasSuffix(name, ".php") &&
		!strings.HasSuffix(name, "_gen.php")
}

func getOutputName(source string) string {
	source, _ = strings.CutSuffix(source, ".php")
	return source + "_gen.php"
}

func generate(sourceFile string) error {
	data, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}
	return os.WriteFile(getOutputName(sourceFile), template.Compile(data), 0o644)
}

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !isSourceFile(entry.Name()) {
			continue
		}

		if err := generate(entry.Name()); err != nil {
			panic(err)
		}
	}
}
