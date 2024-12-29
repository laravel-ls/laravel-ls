
GO = go
VERSION=$(shell git describe --always --tags --dirty --match="v*")
GOLDFLAGS=-s -w -X laravel-ls.Version="$(VERSION)"
GOBUILDFLAGS+=-v -p $(shell nproc) -ldflags="$(GOLDFLAGS)"

.PHONY: build

build:
	mkdir -p build
	$(GO) build $(GOBUILDFLAGS) -o ./build/laravel-ls ./cmd/laravel-ls

test:
	$(GO) test -v ./...
