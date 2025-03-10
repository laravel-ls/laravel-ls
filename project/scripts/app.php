<?php

// Taken from vscode laravel extension.
// https://github.com/laravel/vs-code-extension/blob/main/php-templates/app.php

echo collect(app()->getBindings())
    ->filter(fn ($binding) => ($binding['concrete'] ?? null) !== null)
    ->flatMap(function ($binding, $key) {
        $boundTo = new ReflectionFunction($binding['concrete']);

        $closureClass = $boundTo->getClosureScopeClass();

        if ($closureClass === null) {
            return [];
        }

        return [
            $key => [
                'path' => $closureClass->getFileName(),
                'class' => $closureClass->getName(),
                'line' => $boundTo->getStartLine(),
            ],
        ];
    })->toJson();
