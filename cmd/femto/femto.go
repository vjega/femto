package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/vjega/femto"
)

func saveBuffer(b *femto.Buffer, path string) error {
	return ioutil.WriteFile(path, []byte(b.String()), 0600)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: femto [filename]\n")
		os.Exit(1)
	}
	path := os.Args[1]

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read %v: %v", path, err)
	}

	app := tview.NewApplication()
	buffer := femto.NewBufferFromString(string(content), path)
	root := femto.NewView(buffer)
	root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlS:
			saveBuffer(buffer, path)
			return nil
		case tcell.KeyCtrlQ:
			app.Stop()
			return nil
		}
		return event
	})
	app.SetRoot(root, true)

	if err := app.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
