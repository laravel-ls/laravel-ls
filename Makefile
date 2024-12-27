
GO = go

.PHONY: build

build:
	mkdir -p build
	$(GO) build -o ./build/laravel-ls ./cmd/laravel-ls

test:
	$(GO) test -v ./...
