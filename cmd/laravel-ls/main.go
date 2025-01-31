package main

import (
	"github.com/laravel-ls/laravel-ls/cmd/laravel-ls/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.WithError(err).Error("Error")
	}
}
