; --------------------------------------------------
;  App
; --------------------------------------------------

; app('key')
(function_call_expression
  function: (name) @function (#eq? @function "app")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @app.service
        (encapsed_string . (string_content) .) @app.service
        (encapsed_string "\"" . "\"") @app.service
    ])
  ))


; app()->make('key')
(member_call_expression
  object: (
    function_call_expression
    function: (name) @object.name (#eq? @object.name "app")
    arguments: (arguments "(" . ")"))
  name: (name) @function.name (#eq? @function.name "make" )
  arguments: (arguments
    . (argument [
        (string (string_content)?) @app.service
        (encapsed_string . (string_content) .) @app.service
        (encapsed_string "\"" . "\"") @app.service
    ])
  ))

; App::make('key')
; App::bound('key')
; App::isShared('key')
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class
  ] (#eq? @class "App")
  name: (name) @method
  (#any-of? @method "make" "bound" "isShared")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @app.service
        (encapsed_string . (string_content) .) @app.service
        (encapsed_string "\"" . "\"") @app.service
    ])
  ))

; --------------------------------------------------
;  View
; --------------------------------------------------

; view() calls
(function_call_expression
  function: (name) @function (#eq? @function "view")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @view.name
        (encapsed_string . (string_content) .) @view.name
        (encapsed_string "\"" . "\"") @view.name
    ])
  ))

; Route::view() calls.
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class 
  ] (#eq? @class "Route")
  name: (name) @method (#eq? @method "view")
  arguments: (arguments
    (argument) ; First parameter is the route.
    . (argument [
        (string (string_content)?) @view.name
        (encapsed_string . (string_content) .) @view.name
        (encapsed_string "\"" . "\"") @view.name
    ])
  ))

; --------------------------------------------------
;  Environment
; --------------------------------------------------

; env() calls
(function_call_expression
  function: (name) @function (#eq? @function "env")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @env.key
        (encapsed_string . (string_content) .) @env.key
        (encapsed_string "\"" . "\"") @env.key
    ])
))

; Env::get() calls.
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class 
  ] (#eq? @class "Env")
  name: (name) @method (#eq? @method "get")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @env.key
        (encapsed_string . (string_content) .) @env.key
        (encapsed_string "\"" . "\"") @env.key
    ])
))

; --------------------------------------------------
;  Assets
; --------------------------------------------------

; asset() calls
(function_call_expression
  function: (name) @function (#eq? @function "asset")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @asset.filename
        (encapsed_string (string_content)?) @asset.filename
    ])
  ))

