package main

import (
	"fmt"
	"os"

	"mzarnowski.dev/toolbox/apps/cli"
	"mzarnowski.dev/toolbox/apps/toolbox/profile"
	"mzarnowski.dev/toolbox/config"
)

var commands = map[string]cli.Command{
	"profile": profile.Command,
}

func main() {
	t := config.NewToolbox()

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "No arguments provided")
		os.Exit(1)
	} else if err := run(t, os.Args[1], os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run(toolbox config.Toolbox, name string, args []string) error {
	if command, exists := commands[name]; exists {
		return command(toolbox, args)
	} else {
		return fmt.Errorf("invalid command: %s", name)
	}
}
