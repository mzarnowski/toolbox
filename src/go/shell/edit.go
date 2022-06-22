package shell

import (
	"errors"
	"io"
	"os"

	"mzarnowski.dev/toolbox"
)

func Editor() (string, bool) {
	return os.LookupEnv("EDITOR")
}

func Edit(path string) error {
	editor, isSet := Editor()
	if !isSet {
		return errors.New("editor not set")
	}

	backup, err := toolbox.CachedCopyOfFile(path)
	if err != nil {
		return err
	}

	defer os.Remove(backup.Name())
	if err := Piped(editor, backup.Name()); err != nil {
		return err
	}

	if err := os.Rename(backup.Name(), path); err != nil {
		return err
	}

	return nil
}

func backup(path string) (*os.File, error) {
	source, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	target, err := os.CreateTemp("", "toolbox-edit-*")
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(target, source); err != nil {
		os.Remove(target.Name()) // TODO what about the error?
		return nil, err
	}

	return target, nil
}
