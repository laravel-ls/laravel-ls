package runtime

import (
	"errors"
	"path"

	"github.com/laravel-ls/laravel-ls/utils"
	"github.com/laravel-ls/laravel-ls/utils/repository"
)

var ErrNotAnLaravelProject = errors.New("not a laravel project")

// Environment encapsulates the runtime details needed to
// execute PHP code and retrieve information from a Laravel project.
type Environment struct {
	rootPath string
	process  *PHPProccess
}

// NewEnvironment initializes a new environment by
// analyzing the project in rootPath and finding a suitable php process.
func NewEnvironment(rootPath string) (*Environment, error) {
	if utils.FileExists(path.Join(rootPath, "bootstrap", "app.php")) {
		return nil, ErrNotAnLaravelProject
	}

	process, err := FindPHPProcess(rootPath)
	if err != nil {
		return nil, err
	}

	return &Environment{
		rootPath: rootPath,
		process:  process,
	}, nil
}

// RootPath returns the root path for the environment
func (env Environment) RootPath() string {
	return env.rootPath
}

// RootPath returns the php process for the environment
func (env Environment) Process() *PHPProccess {
	return env.process
}

// Configs retrieves the configuration repository in the environment
func (env Environment) Configs() (repository.ConfigRepository, error) {
	return GetConfigs(env.process, env.rootPath)
}

// AppBindings retrieves the application bindings repository in the environment
func (env Environment) AppBindings() (repository.AppRepository, error) {
	return GetAppBindings(env.process, env.rootPath)
}
