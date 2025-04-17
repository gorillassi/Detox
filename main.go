package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"

	"soulblog/cli/commands"
)

func main() {
	app := &cli.App{
		Name:  "detox",
		Usage: "P2P Onion Blog CLI",
		Commands: []*cli.Command{
			commands.Init,
			commands.Post,
			commands.Feed,
			commands.Sync,
			commands.Exit,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}