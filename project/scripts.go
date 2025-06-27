package project

import _ "embed"

//go:embed scripts/app_gen.php
var appScript []byte

//go:embed scripts/configs_gen.php
var configScript []byte

//go:embed scripts/routes_gen.php
var routeScript []byte
