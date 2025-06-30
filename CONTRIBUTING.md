# Contributing to Laravel Language Server

Thank you for your interest in contributing to Laravel Language Server! 
This project aims to provide intelligent code assistance for Laravel applications using the Language Server Protocol (LSP).

## Requirements

To contribute and run the project locally, you’ll need the following tools installed:

### Core Tools

* [Go 1.23+](https://go.dev/doc/install) - The main language for the language server
* [PHP 8.1+](https://www.php.net/downloads) - Required to analyze and introspect Laravel applications
* **Make** - Used to run common development tasks

  * **Linux/macOS:** Preinstalled (on most systems, otherwise consult your distros documentation)
  * **Windows:** Use via WSL, Git Bash, or install from [GNUWin](http://gnuwin32.sourceforge.net/packages/make.htm)

* [GolangCI-Lint](https://golangci-lint.run/usage/install/) - Linter for Go code

  Install via:

  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

* **C Compiler (for cgo)**
  Required to compile native bindings (e.g., Tree-sitter parsers)

  Laravel Language Server uses [Tree-sitter](https://tree-sitter.github.io/tree-sitter/) via cgo for fast and accurate parsing.

  * **Linux:** Install `gcc` or `clang`
    Example: `sudo apt install build-essential`
  * **macOS:** Install Xcode Command Line Tools

    ```bash
    xcode-select --install
    ```
  * **Windows:** Install `gcc` via [MSYS2](https://www.msys2.org/) or use WSL with a Linux toolchain

  Development is mainly done in linux with gcc, so any other os/compiler combination may not work.
  Feel free to open a issue if you encounter some difficulties compiling the project and we can sort it out.

## Getting Started

1. **Clone the repository:**

```bash
git clone https://github.com/laravel-ls/laravel-ls.git
cd laravel-ls
```

2. **Build the language server:**

```bash
make build
```

3. **Configure your editor:**

See https://github.com/laravel-ls/laravel-ls?tab=readme-ov-file#configure

And modify the path to your local binary in `build/laravel-ls`

TIP: On linux, you can use the `start.sh` bash script that will compile and run the server on every start/restart.

## Architecture

The Laravel Language Server is structured to be modular, efficient, and Laravel-aware. 
Below is an overview of key components in the codebase

### Plugin System (providers)

The plugin system is located in the `provider` package and contains the interfaces that providers can implement.

Concrete providers are Located in `laravel/providers`, this system allows different Laravel-aware features 
(like `view()`, `config()`, or `env()` support) to hook into the LSP life cycle. 
Each plugin can implement capabilities such as diagnostics, completions, hover, and go-to definition.

Providers are registered during server startup and dispatched based on file type and symbol usage.

### Parsing (Tree-sitter)

Parsing source files is a central part of an analysis tool like an LSP server and laravel-ls uses [Tree-sitter](https://tree-sitter.github.io/) for this.
Tree-sitter enables fast, incremental parsing of source code suitable for an editor.

The `parser` package integrates treesitter via `cgo` to parse PHP code efficiently.
Tree-sitter grammars are compiled as C libraries and accessed directly in Go through native bindings.

### Static vs Dynamic Analysis

Some features, like autocomplete for `app('...')` or `config('...')` rely on understanding the structure and contents of Laravel files.
While basic resolution can be done via static analysis (e.g., AST parsing with Tree-sitter), many Laravel features are 
best understood through **dynamic evaluation**, this means that laravel-ls runs the laravel application to gather **runtime** information.

This is implemented in two main packages:

- `runtime` - responsible for setting up a correct environment to execute php code.
- `project` - php scripts that gather the information used by the providers.

### LSP

The language server itself is implemented in the `lsp` package and takes care of the protocol part, delegating tasks
to the appropriate `provider`

## Contributing Guidelines

* Format code using `go fmt ./...`
* Run `make lint` before submitting a PR
* Use meaningful commit messages
* Don't commit everything into one large commit. Keep them atomic and focused

Feel free to open issues, suggest features, or submit pull requests!

