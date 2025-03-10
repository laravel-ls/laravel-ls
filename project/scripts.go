package project

import _ "embed"

//go:embed scripts/app_gen.php
var appScript []byte

//go:embed scripts/configs_gen.php
var configScript []byte
