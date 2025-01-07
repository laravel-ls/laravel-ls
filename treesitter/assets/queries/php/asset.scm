; asset() calls
(function_call_expression
  function: (name) @function (#eq? @function "asset")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @asset.filename
        (encapsed_string (string_content)?) @asset.filename
    ])
))

