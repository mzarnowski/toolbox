package main

import (
	"fmt"
	"os"
	"os/exec"

	"mzarnowski.dev/toolbox/paths"
)

const entrypoint string = "apps/toolbox/main.go"

func main() {
	binary := paths.BinaryPath("toolbox")
	cmd := exec.Command("go", "build", "-o", binary, entrypoint)
	cmd.Dir = paths.ConfigPath("toolbox", "src/go")
	cmd.Stderr = os.Stderr

	fmt.Println(cmd)
	if err := cmd.Run(); err != nil {
		panic(err.Error())
	}
}
