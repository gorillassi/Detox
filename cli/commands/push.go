package commands

import (
	"soulblog/crypto"
	"fmt"
	"os"
	"sort"
	"encoding/json"
	"soulblog/core"
	"soulblog/p2p"
	"path/filepath"
	"soulblog/storage"
	"github.com/urfave/cli/v2"
)

var Push = &cli.Command{
	Name:  "push",
	Usage: "publish your latest post to the P2P network",
	Action: func(c *cli.Context) error {
		dir := storage.ResolvePath(StorageDir)

		_, err := crypto.LoadKeys(dir)
		if err != nil {
			return fmt.Errorf("failed to load keys: %w", err)
		}

		postDir := filepath.Join(dir, "posts")
		files, err := os.ReadDir(postDir)
		if err != nil || len(files) == 0 {
			return fmt.Errorf("no posts to push")
		}

		sort.Slice(files, func(i, j int) bool {
			infoI, errI := files[i].Info()
			infoJ, errJ := files[j].Info()
			if errI != nil || errJ != nil {
				return false
			}
			return infoI.ModTime().After(infoJ.ModTime())
		})

		latest := files[0]
		data, err := os.ReadFile(filepath.Join(postDir, latest.Name()))
		if err != nil {
			return err
		}

		var post core.Post
		if err := json.Unmarshal(data, &post); err != nil {
			return fmt.Errorf("invalid post format: %w", err)
		}

		if err := p2p.PublishPost(data); err != nil {
			return fmt.Errorf("failed to publish post: %w", err)
		}

		fmt.Println("[+] Post published:", latest.Name())
		return nil
	},
}