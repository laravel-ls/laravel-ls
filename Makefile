
GO = go
VERSION=$(shell git describe --always --tags --dirty --match="v*")
GOLDFLAGS=-s -w -X main.version="$(VERSION)"


ifeq ($(OS), Windows_NT)
	PROGRAM=./build/laravel-ls.exe
	GOBUILDFLAGS+=-v -ldflags="$(GOLDFLAGS)"
	MKBUILDDIR=
else ifeq ($(shell uname -s), Darwin)
	PROGRAM=./build/laravel-ls
	GOBUILDFLAGS+=-v -p $(shell sysctl -n hw.logicalcpu) -ldflags="$(GOLDFLAGS)"
	MKBUILDDIR=mkdir -p ./build
else
	PROGRAM=./build/laravel-ls
	GOBUILDFLAGS+=-v -p $(shell nproc) -ldflags="$(GOLDFLAGS)"
	MKBUILDDIR=mkdir -p ./build
endif

.PHONY: build

build:
	$(MKBUILDDIR)
	$(GO) build $(GOBUILDFLAGS) -o $(PROGRAM) ./cmd/laravel-ls

generate:
	$(GO) generate ./...

lint:
	golangci-lint run ./...

test:
	$(GO) test -v ./...
