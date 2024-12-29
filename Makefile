
GO = go
GOLDFLAGS=-s -w
GOBUILDFLAGS+=-v -p $(shell nproc) -ldflags="$(GOLDFLAGS)"

.PHONY: build

build:
	mkdir -p build
	$(GO) build $(GOBUILDFLAGS) -o ./build/laravel-ls ./cmd/laravel-ls

test:
	$(GO) test -v ./...
