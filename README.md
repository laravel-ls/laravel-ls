[![Test](https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml/badge.svg)](https://github.com/laravel-ls/laravel-ls/actions/workflows/test.yml)

# Laravel-ls

Laravel Language Server written in go.

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

- [ ] Hover information sow path to the component
- [ ] Auto-complete existing components
- [ ] Auto-complete arguments defined in the component 
- [ ] Diagnostics for components that do not exists.

## Other features on the horizon

- Livewire support
- Inertia support
- Eloquent support
- Jump to test file from class.

## Author

Henrik Hautakoski <henrik@shufflingpixels.com>
