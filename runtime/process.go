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

// Exec executes the PHP code in the specified working directory.
// It returns an io.Reader with the output or an error if execution fails.
func (proc PHPProcess) Exec(workingDir string, code []byte) (io.Reader, error) {
	outBuf := &bytes.Buffer{}
	errBuf := &strings.Builder{}

	// hash code content for temporary file name
	hash := md5.New()
	hash.Write(code)
	hashValue := hash.Sum(nil)
	filePath := path.Join(os.TempDir(), fmt.Sprintf("laravel-ls-%x.php", hashValue))

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

	cmd := exec.Command(proc.Args[0], append(proc.Args[1:], filePath)...)
	cmd.Dir = workingDir
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	if cmdErr := cmd.Run(); cmdErr != nil {
		err := errors.New(errBuf.String())
		return nil, errors.Join(err, cmdErr)
	}
	return outBuf, nil
}
