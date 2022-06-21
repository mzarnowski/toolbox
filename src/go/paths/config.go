package paths

import "fmt"

func ConfigHome() string { return getOrElse("XDG_CONFIG_HOME", "$HOME/.config") }

func ConfigPath(application, path string) string {
	return fmt.Sprintf("%s/%s/%s", ConfigHome(), application, path)
}
