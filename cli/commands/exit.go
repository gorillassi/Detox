// cli/commands/exit.go
package commands

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
)

var Exit = &cli.Command{
	Name:  "exit",
	Usage: "exit the application",
	Action: func(c *cli.Context) error {
		fmt.Println("Exiting...")
		os.Exit(0)
		return nil
	},
}
