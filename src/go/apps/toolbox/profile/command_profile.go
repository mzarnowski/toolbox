package profile

import (
	"fmt"
	"strings"

	"mzarnowski.dev/toolbox/apps/cli"
	"mzarnowski.dev/toolbox/config"
)

var commands = map[string]cli.Command{
	"edit": edit,
}

func Command(toolbox config.Toolbox, args []string) error {
	if len(args) == 0 {
		fmt.Println(toolbox.DefaultProfile())
		return nil
	} else if !strings.Contains(args[0], "=") {
		command, _ := commands[args[0]]
		return command(toolbox, args[1:])
	}
	return nil
}
