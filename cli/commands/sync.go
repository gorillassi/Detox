// cli/commands/sync.go
package commands

import (
	cli "github.com/urfave/cli/v2"
	"soulblog/p2p"
	"soulblog/storage"
)

var Sync = &cli.Command{
	Name:  "sync",
	Usage: "start syncing with peers",
	Action: func(c *cli.Context) error {
		dir := storage.ResolvePath(StorageDir)
		return p2p.StartSync(dir)
	},
}
