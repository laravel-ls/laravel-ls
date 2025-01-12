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
