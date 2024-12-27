package transport

import (
	"errors"
	"os"
)

// Needed to implement io.ReadWriteCloser interface
//
//	type stdio struct {
//		in  io.ReadCloser
//		out io.WriteCloser
//	}
//
//	func NewStdio(in io.ReadCloser, out io.WriteCloser) stdio {
//		return stdio{
//			in:  in,
//			out: out,
//		}
//	}
type stdio struct{}

func NewStdio() stdio {
	return stdio{}
}

func (s stdio) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (s stdio) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (s stdio) Close() error {
	var errs []error
	if err := os.Stdout.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := os.Stdin.Close(); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}
