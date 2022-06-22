package paths

import "fmt"

func CacheRoot() string { return getOrElse("XDG_CONFIG_HOME", "$HOME/.cache") }

func CacheHome(application string) string {
	return fmt.Sprintf("%s/%s", CacheRoot(), application)
}

func CachePath(application, path string) string {
	return fmt.Sprintf("%s/%s/%s", CacheRoot(), application, path)
}
