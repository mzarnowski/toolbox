package main

import (
	"fmt"
	"io/fs"
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

	directories := []string{
		paths.CacheHome("toolbox"),
		paths.ConfigHome("toolbox"),
	}

	for _, dir := range directories {
		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			fail(err)
		}
	}

	if err := cmd.Run(); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
