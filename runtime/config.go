package runtime

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/utils/repository"
)

//go:embed scripts/configs_gen.php
var configScript []byte

// GetConfigs retrieves the configuration repository in rootPath by using the provided php process.
func GetConfigs(call *PHPProccess, rootPath string) (repository.ConfigRepository, error) {
	return CallScript(call, rootPath, configScript, repository.ConfigRepository{})
}
