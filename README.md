<h1 align="center">Laravel-ls</h1>

<p align="center">
    Laravel Language Server written in go.
    <br />
    <a href="#about">About</a>
    |
    <a href="#features">Features</a>
    |
    <a href="#install">Install</a>
    |
    <a href="#build">Building</a>
</p>

<p align="center">
    <a href="https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml">
        <img src="https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml/badge.svg" />
    </a>
</p>

## About

Laravel-ls is a tool that enhances your text editor or IDE with
powerful features specifically designed for [Laravel](https://laravel.com) projects.

By implementing the [Language Server Protocol (LSP)](https://microsoft.github.io/language-server-protocol/),
Laravel-ls seamlessly integrates with any LSP-compatible editor, bringing advanced capabilities
such as intelligent auto-completion, navigation to definitions, real-time diagnostics, and more.

Although still in its early development stages, Laravel-ls aims to provide a
streamlined and efficient development experience tailored to the Laravel framework.

## Features

List of done and planned features of the language server.

### Views

```php
view('home')
Route::view('/user/profile', 'user.profile')
```

- [x] Hover information shows view template path.
- [x] Go to definition opens the template
- [x] Auto-complete existing template names
- [ ] Auto-complete for existing arguments present in the template
- [x] Diagnostics for template files that do not exists
- [ ] Code action to create view file that do not exists.

### Environment

```php
env('APP_NAME');
Env::get('APP_NAME');
```

- [x] Hover information shows actual value from `.env` file
- [x] Go to `.env` file and key location
- [x] Auto-complete for defined keys.
- [x] Diagnostics for non defined keys.
- [x] Code action to create missing key

### Config

```php
config('app.name')
config()->string('app.name')
config()->getMany(['app.name'])
Config::get('app.name')
Config::getMany(['app.name'])
```

- [x] Hover information shows actual value.
- [x] Auto-complete for defined config keys
- [x] Go to config file and value location from key
- [x] Diagnostics for non defined config keys.

### Application bindings

```php
app('db.connection')
app()->make('db.connection')
app()->bound('db.connection')
app()->isShared('db.connection')
app()->make('db.connection')
app()->bound('db.connection')
app()->isShared('db.connection')
```

- [x] Hover information shows value.
- [x] Auto-complete for defined bindings
- [x] Go to location where binding is defined
- [x] Diagnostics for non defined bindings

### Assets

```php
asset('main.css');
```

- [x] Auto-completion
- [x] Go to asset file
- [x] Diagnostics for non existent assets

### Blade components

```html
<x-component.name />
<x-component name="dynamic-component" />
```

- [ ] Hover information show path to the component
- [ ] Auto-complete existing components
- [ ] Auto-complete arguments defined in the component
- [ ] Diagnostics for components that do not exists.
- [ ] Code action to create missing components

## Other features on the horizon

- Livewire support
- Inertia support
- Eloquent support
- Jump to test file from class.

## Install

The program can be installed via go:

```sh
go install github.com/laravel-ls/laravel-ls/cmd/laravel-ls@latest
```

And if you have added [GOPATH](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable) to your shell's `PATH`. You should be able to just run the server with:

```sh
laravel-ls
```

See the official documentation of [go install](https://go.dev/ref/mod#go-install)

## Build

To build the project you need golang version `1.22` or later, `make` and a c compiler.

When the dependencies are met, running `make` will compile and produce the
binary `build/laravel-ls`.

## Configure

### Neovim

The LSP server can be started like any other server via `vim.lsp.start` and an auto-command.

Just change the path to the correct directory on your filesystem

```lua
vim.api.nvim_create_autocmd("FileType", {
    pattern = { "php", "blade" },
    callback = function ()
        vim.lsp.start({
            name = "laravel-ls",
            cmd = { '/path/to/laravel-ls/build/laravel-ls' },
            -- if you want to recompile everytime
            -- the language server is started.
            -- Uncomment this line instead
            -- cmd = { '/path/to/laravel-ls/start.sh' },
            root_dir = vim.fn.getcwd(),
        })
    end
})
```

## Author

Henrik Hautakoski <henrik@shufflingpixels.com>
