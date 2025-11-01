GO = go
VERSION=$(shell git describe --always --tags --dirty --match="v*")
GOLDFLAGS=-s -w -X main.version="$(VERSION)"
GOBUILDFLAGS += -v -ldflags="$(GOLDFLAGS)"

ifeq ($(OS), Windows_NT)
	PROGRAM=./build/laravel-ls.exe
	MKBUILDDIR=
# Uncomment this when macos setup diverges from linux.
# else ifeq ($(shell uname -s), Darwin)
# 	PROGRAM=./build/laravel-ls
# 	MKBUILDDIR=mkdir -p ./build
else
	PROGRAM=./build/laravel-ls
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
