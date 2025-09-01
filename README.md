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

<div align="center">

![Go Version](https://img.shields.io/github/go-mod/go-version/laravel-ls/laravel-ls?style=for-the-badge)
[![License](https://img.shields.io/github/license/laravel-ls/laravel-ls?style=for-the-badge)](/LICENSE)
[![Test](https://img.shields.io/github/actions/workflow/status/laravel-ls/laravel-ls/test.yml?style=for-the-badge&label=test)](https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/laravel-ls/laravel-ls?include_prereleases&style=for-the-badge)](https://github.com/laravel-ls/laravel-ls/releases/latest)

</div>

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

### Routes

```php
route('dashboard');
redirect()->route('dashboard');
URL::route('dashboard');
```

- [ ] Hover information shows route definition
- [ ] Go to definition
- [ ] Auto-complete existing route names
- [ ] Auto-complete route arguments
- [ ] Diagnostics for routes that do not exists
- [ ] Code action to create a new route if route do not exists.

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

###  Download via github

Official binaries for Windows and Linux are provided on each [github release](https://github.com/laravel-ls/laravel-ls/releases)

MacOS users have to use [Install via go](#install-via-go) 

> NOTE: although MacOS is not officially supported, some users have had success building it.

Just download the program and make sure its located somewhere in your `$PATH`

Example command to download:
```sh
sudo wget -O /usr/local/bin/laravel-ls https://github.com/laravel-ls/laravel-ls/releases/download/<VERSION>/laravel-ls-<VERSION>-linux-amd64 & \
   sudo chmod 755 /usr/local/bin/laravel-ls
```

### Install via go

The program can be compiled and installed via go:

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

The project officially supports building Windows and Linux binaries. But some users have had success on MacOS.

When the dependencies are met, running `make` will compile and produce the
binary `build/laravel-ls`.

## Configure

### Neovim

#### nvim-lspconfig

##### nvim < 0.11

```lua
require'lspconfig'.laravel_ls.setup{}
```

or custom config 

```lua
require'lspconfig'.laravel_ls.setup{
    -- Server-specific settings. See `:help lspconfig-setup`
    settings = {
        cmd = { …  },
    },
}
```

##### nvim => 0.11

```lua
vim.lsp.enable('laravel_ls')
```

or custom config

```lua
vim.lsp.config('laravel_ls', {
    cmd = { …  },
})
```

All settings can be found [here](https://github.com/neovim/nvim-lspconfig/blob/master/doc/configs.md#laravel_ls)

#### native

The LSP server can be started like any other server via `vim.lsp.start` and an auto-command.

Just change the path to the correct directory on your filesystem

```lua
vim.api.nvim_create_autocmd("FileType", {
    pattern = { "php", "blade" },
    callback = function ()
        vim.lsp.start({
            name = "laravel-ls",

            -- if laravel ls is in your $PATH
            cmd = { 'laravel-ls' },
            
            -- Absolute path
            -- cmd = { '/path/to/laravel-ls/build/laravel-ls' },
            
            -- if you want to recompile everytime
            -- the language server is started.
            -- cmd = { '/path/to/laravel-ls/start.sh' },

            root_dir = vim.fn.getcwd(),
        })
    end
})
```


## Author

Henrik Hautakoski <henrik@shufflingpixels.com>
