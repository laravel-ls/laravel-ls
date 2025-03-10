
GO = go
VERSION=$(shell git describe --always --tags --dirty --match="v*")
GOLDFLAGS=-s -w -X main.version="$(VERSION)"
GOBUILDFLAGS+=-v -p $(shell nproc) -ldflags="$(GOLDFLAGS)"

.PHONY: build

build:
	mkdir -p build
	$(GO) build $(GOBUILDFLAGS) -o ./build/laravel-ls ./cmd/laravel-ls

generate: 
	$(GO) generate ./...

lint:
	golangci-lint run ./...

test:
	$(GO) test -v ./...
