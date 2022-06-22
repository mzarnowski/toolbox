package paths

import "fmt"

func ConfigRoot() string { return getOrElse("XDG_CONFIG_HOME", "$HOME/.config") }

func ConfigHome(application string) string {
	return fmt.Sprintf("%s/%s", ConfigRoot(), application)
}

func ConfigPath(application, path string) string {
	return fmt.Sprintf("%s/%s/%s", ConfigRoot(), application, path)
}
