package main

import (
	"bytes"
	"fmt"

	"mzarnowski.dev/toolbox/cli"
)

func main() {
	str := "a\tb\tc\nd\tea\tf"
	var reader = bytes.NewReader([]byte(str))
	table := cli.NewTable([]string{"foo", "bar", "baz"}, cli.NewRuneIterator(reader))
	for _, header := range table.Header {
		fmt.Printf("%s\t", header)
	}
	for {
		if row, err := table.Next(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		} else {
			fmt.Println()
			for _, value := range row {
				fmt.Printf("%s\t", value)
			}
		}
	}
}
