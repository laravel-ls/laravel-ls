package provider

import (
	"github.com/laravel-ls/laravel-ls/cache"
	"github.com/laravel-ls/laravel-ls/parser"

	log "github.com/sirupsen/logrus"
)

type BaseContext struct {
	Logger    *log.Entry
	File      *parser.File
	FileCache *cache.FileCache
}
