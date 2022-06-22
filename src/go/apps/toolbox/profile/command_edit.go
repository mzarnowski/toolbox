package profile

import (
	"mzarnowski.dev/toolbox"
	"mzarnowski.dev/toolbox/config"
	"mzarnowski.dev/toolbox/shell"
)

func edit(t config.Toolbox, args []string) error {
	return shell.Edit(toolbox.LocalConfigFile())
}
