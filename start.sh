#!/bin/bash
# Simple script to compile and run the server (useful when developing)
cd $(dirname $(readlink -f $BASH_SOURCE)); go run ./cmd/laravel-ls/main.go
