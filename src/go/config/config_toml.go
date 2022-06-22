package config

import (
	"errors"
	"os"

	"mzarnowski.dev/toolbox"
	"mzarnowski.dev/toolbox/encoding"
)

const PROFILE_NONE = ""

type tomlToolbox struct {
	parsed map[string]any
}

func NewToolbox() Toolbox {
	receiver := map[string]any{
		"Default-Profile": PROFILE_NONE,
	}

	configPath := toolbox.LocalConfigFile()
	configFile := encoding.NewTomlFile[map[string]any](configPath)

	err := configFile.Read(&receiver)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err.Error())
	} else {
		return &tomlToolbox{receiver}
	}
}

func (t *tomlToolbox) DefaultProfile() string {
	return t.parsed["Default-Profile"].(string)
}
