package provider

import (
	"laravel-ls/cache"
	"laravel-ls/parser"

	log "github.com/sirupsen/logrus"
)

type BaseContext struct {
	Logger    *log.Entry
	File      *parser.File
	FileCache *cache.FileCache
}
