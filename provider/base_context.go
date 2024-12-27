package provider

import (
	"github.com/shufflingpixels/laravel-ls/cache"
	"github.com/shufflingpixels/laravel-ls/parser"
	log "github.com/sirupsen/logrus"
)

type BaseContext struct {
	Logger    *log.Entry
	File      *parser.File
	FileCache *cache.FileCache
}
