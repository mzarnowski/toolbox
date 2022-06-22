package toolbox

import (
	"errors"
	"io"
	"os"

	"github.com/google/uuid"
	"mzarnowski.dev/toolbox/paths"
)

func LocalConfigFile() string {
	return paths.ConfigPath("toolbox", "config.toml")
}

func CachedCopyOfFile(path string) (*os.File, error) {
	id := uuid.NewString()

	target, err := os.Create(paths.CachePath("toolbox", id))
	if err != nil {
		return nil, err
	}

	source, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return target, nil
	} else if err != nil {
		return nil, err
	}

	if _, err := io.Copy(target, source); err != nil {
		os.Remove(target.Name())
		return nil, err
	} else {
		return target, nil
	}
}
