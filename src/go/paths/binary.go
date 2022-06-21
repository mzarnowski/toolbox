package paths

import (
	"fmt"
	"os"
)

func BinaryPath(application string) string {
	return fmt.Sprintf("%s/%s", os.ExpandEnv("${HOME}/.local/bin"), application)
}
