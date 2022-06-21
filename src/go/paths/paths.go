package paths

import (
	"os"
)

func getOrElse(name, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	} else {
		return os.ExpandEnv(fallback)
	}
}
