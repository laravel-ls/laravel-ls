; env() calls
(function_call_expression
  function: (name) @function (#eq? @function "env")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @env.key
        (encapsed_string (string_content)?) @env.key
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
        (encapsed_string (string_content)?) @env.key
    ])
))
