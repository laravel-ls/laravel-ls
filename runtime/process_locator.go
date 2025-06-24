package runtime

import (
	"errors"
	"os/exec"
	"path"
)

var ErrNoBinary = errors.New("failed to find php binary")

// PHPProcessLocator is a function type that returns a pointer to a PHPProcess.
// It's used to locate a valid PHP executable in different environments.
type PHPProcessLocator func() *PHPProcess

// sail checks if the Laravel Sail binary is available
// (and sail is running) and returns a PHPProcessLocator that can be used to execute code inside sail.
func sail(rootPath string) PHPProcessLocator {
	sailBinary := path.Join(rootPath, "vendor/bin/sail")
	return func() *PHPProcess {
		// Check if sail is running
		if err := exec.Command(sailBinary, "ps").Run(); err == nil {
			return NewPHPProcess(sailBinary, "php", "-r")
		}
		return nil
	}
}

// local searches for the system PHP binary in the system's PATH.
func local() *PHPProcess {
	p, err := exec.LookPath("php")
	if err != nil {
		return nil
	}
	return NewPHPProcess(p, "-r")
}

func FindPHPProcess(rootPath string) (*PHPProcess, error) {
	locators := []PHPProcessLocator{
		sail(rootPath),
		local,
	}

	for _, locator := range locators {
		if c := locator(); c != nil {
			return c, nil
		}
	}

	return nil, ErrNoBinary
}
