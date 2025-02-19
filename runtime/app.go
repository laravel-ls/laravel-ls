package runtime

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/utils/repository"
)

//go:embed scripts/app_gen.php
var appScript []byte

// GetAppBindings retrieves the application bindings
// repository in rootPath by using the provided php process.
func GetAppBindings(call *PHPProccess, rootPath string) (repository.AppRepository, error) {
	return CallScript(call, rootPath, appScript, repository.AppRepository{})
}
