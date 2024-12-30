[![Test](https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml/badge.svg)](https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml)

<h1>Laravel-ls</h1>
<p align="center">
    Laravel Language Server written in go.
    <br />
    <a href="#about">About</a>
    |
    <a href="#features">Features</a>
    |
    <a href="#build">Building</a>
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

### Environment

```php
env('APP_NAME');
Env::get('APP_NAME');
```

- [x] Hover information shows actual value from `.env` file
- [x] Go to `.env` file and key location
- [x] Auto-complete for defined keys.
- [x] Diagnostics for non defined keys.
- [ ] Code action to create missing key

### Config

```php
config('app.name')
Config::get('app.name')
```

- [ ] Hover information shows actual value.
- [ ] Auto-complete for defined config keys
- [ ] Go to config file and value location from key
- [ ] Diagnostics for non defined config keys.

### Blade components

```html
<x-component.name />
<x-component name="dynamic-component" />
```

- [ ] Hover information show path to the component
- [ ] Auto-complete existing components
- [ ] Auto-complete arguments defined in the component
- [ ] Diagnostics for components that do not exists.

## Other features on the horizon

- Livewire support
- Inertia support
- Eloquent support
- Jump to test file from class.

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
