package genny

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// ForceFile is a TransformerFn that will return an error if the path exists if `force` is false. If `force` is true it will delete the path.
func ForceFile(f File, force bool) TransformerFn {
	return func(f File) (File, error) {
		path := f.Name()
		path, err := filepath.Abs(path)
		if err != nil {
			return f, errors.WithStack(err)
		}
		_, err = os.Stat(path)
		if err != nil {
			// path doesn't exist. move on.
			return f, nil
		}
		if !force {
			return f, errors.Errorf("path %s already exists", path)
		}
		if err := os.RemoveAll(path); err != nil {
			return f, errors.WithStack(err)
		}
		return f, nil
	}
}

// Force is a RunFn that will return an error if the path exists if `force` is false. If `force` is true it will delete the path.
// Is is recommended to use ForceFile when you can.
func Force(path string, force bool) RunFn {
	if path == "." || path == "" {
		pwd, _ := os.Getwd()
		path = pwd
	}
	return func(r *Runner) error {
		path, err := filepath.Abs(path)
		if err != nil {
			return errors.WithStack(err)
		}
		_, err = os.Stat(path)
		if err != nil {
			// path doesn't exist. move on.
			return nil
		}
		if !force {
			return errors.Errorf("path %s already exists", path)
		}
		if err := os.RemoveAll(path); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}
