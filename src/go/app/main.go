package main

import (
	"bytes"
	"time"

	"github.com/rivo/tview"
	"mzarnowski.dev/toolbox/cli"
)

func main() {
	str := "a\tb\tc\nd\tea\tf"
	var reader = bytes.NewReader([]byte(str))
	table := cli.NewTable([]string{"foo", "bar", "baz"}, cli.NewRuneIterator(reader))

	foo := cli.NewFoo(table.Header, 5)

	app := tview.NewApplication().SetRoot(foo.Table, true)
	channel := make(chan error)

	go run(app, channel)
	if row, err := table.Next(); err == nil {
		d, _ := time.ParseDuration("1s")
		app.Draw().QueueUpdateDraw(foo.Insert(3, row))
		time.Sleep(d)
		app.Draw().QueueUpdateDraw(foo.Insert(1, row))
	}

	if err := <-channel; err != nil {
		panic(err)
	}
}

func run(app *tview.Application, output chan error) {
	if err := app.Run(); err != nil {
		output <- err
	}
	close(output)
}
