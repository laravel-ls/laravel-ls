package main

import (
	"github.com/laravel-ls/laravel-ls/cmd/laravel-ls/cmd"
	"github.com/laravel-ls/laravel-ls/program"
)

var version string = ""

func main() {
	program.VersionOverride = version

	cmd.Run()
}
