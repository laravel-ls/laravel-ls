package runtime

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Process interface {
	Exec(workingDir string, code []byte) (io.Reader, error)
}

// PHPProcess represents a PHP execution process with configurable arguments.
type PHPProcess struct {
	Args []string
}

// NewPHPProcess creates a new PHPProcess instance with the given arguments.
func NewPHPProcess(args ...string) *PHPProcess {
	return &PHPProcess{
		Args: args,
	}
}

func (proc PHPProcess) vendorDir(workingDir string) (string, error) {
	tmpDir := path.Join(workingDir, "vendor", "_laravel-ls")
	if err := os.MkdirAll(tmpDir, 0o755); err != nil {
		return "", err
	}
	return tmpDir, nil
}

// Exec executes the PHP code in the specified working directory.
// It returns an io.Reader with the output or an error if execution fails.
func (proc PHPProcess) Exec(workingDir string, code []byte) (io.Reader, error) {
	outBuf := &bytes.Buffer{}
	errBuf := &strings.Builder{}

	vendorDir, err := proc.vendorDir(workingDir)
	if err != nil {
		return nil, err
	}
	
	// hash code content for temporary file name
	filePath := path.Join(vendorDir, fmt.Sprintf("laravel-ls-%x.php", md5.Sum(code)))

	// Check if the temporary file already exists
	if _, err := os.Stat(filePath); err != nil {
		// If the file does not exist, create it and write the code to it.
		f, err := os.Create(filePath)
		if err != nil {
			return nil, errors.New("failed to create temporary file for PHP code: " + err.Error())
		}

		if _, err := f.Write(code); err == nil {
			f.Close()
		}
	}


	// Get the file relative to the working directory
	relFilePath, _ := filepath.Rel(workingDir, filePath)

	cmd := exec.Command(proc.Args[0], append(proc.Args[1:], relFilePath)...)
	cmd.Dir = workingDir
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	if cmdErr := cmd.Run(); cmdErr != nil {
		err := errors.New(errBuf.String())
		return nil, errors.Join(err, cmdErr)
	}
	return outBuf, nil
}
