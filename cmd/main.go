package main

import (
	"fmt"
	"os"

	"github.com/shalldie/tnote/gist"
	"github.com/shalldie/tnote/note"
)

var (
	FetchOptions gist.FetchOptions
)

func main() {

	token := os.Getenv("TNOTE_GIST_TOKEN")

	if token == "" {
		fmt.Println("Can't find $TNOTE_GIST_TOKEN in $PATH")
		os.Exit(1)
	}

	note.NewTNote(token).Setup()
}
