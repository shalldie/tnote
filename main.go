package main

import (
	"fmt"
	"os"

	_ "github.com/shalldie/tnote/internal/conf"

	"github.com/shalldie/tnote/internal/app"
)

func main() {

	token := os.Getenv("TNOTE_GIST_TOKEN")

	if token == "" {
		fmt.Println("Can't find $TNOTE_GIST_TOKEN in $PATH")
		os.Exit(1)
	}

	app.Run(token)

}
