package cli

import "mzarnowski.dev/toolbox/config"

type Command = func(config.Toolbox, []string) error
