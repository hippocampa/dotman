package main

import (
	"log"
	"os"

	"github.com/hippocampa/dotshadow/cmd"
	"github.com/hippocampa/dotshadow/cmd/subcommand"
)

func main() {
	app := cmd.New()
	app.AddSubCmd(subcommand.InitCmd())
	if err := app.CliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
