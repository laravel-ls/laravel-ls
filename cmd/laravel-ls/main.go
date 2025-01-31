package main

import (
	"github.com/laravel-ls/laravel-ls/cmd/laravel-ls/cmd"
	"github.com/laravel-ls/laravel-ls/program"
	log "github.com/sirupsen/logrus"
)

var version string = ""

func main() {
	program.VersionOverride = version

	if err := cmd.Run(); err != nil {
		log.WithError(err).Error("Error")
	}
}
