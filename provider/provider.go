package provider

import (
	"github.com/shufflingpixels/laravel-ls/cache"
	log "github.com/sirupsen/logrus"
)

type InitContext struct {
	Logger    *log.Entry
	RootPath  string
	FileCache *cache.FileCache
}

type Provider interface {
	Register(manager *Manager)
	Init(ctx InitContext)
}
