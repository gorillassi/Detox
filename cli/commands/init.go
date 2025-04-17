// cli/commands/init.go
package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
	"soulblog/crypto"
	"soulblog/storage"
)

const StorageDir = "~/.soulblog"

var Init = &cli.Command{
	Name:  "init",
	Usage: "generate keypair and init storage",
	Action: func(c *cli.Context) error {
		dir := storage.ResolvePath(StorageDir)
		if err := storage.EnsureDir(dir); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		_, err := crypto.GenerateAndSaveKeys(dir)
		if err != nil {
			return fmt.Errorf("failed to generate/save keys: %w", err)
		}
		fmt.Println("[+] Keys initialized at", dir)
		return nil
	},
}
