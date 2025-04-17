// cli/commands/feed.go
package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	cli "github.com/urfave/cli/v2"
	"soulblog/core"
	"soulblog/storage"
)

var Feed = &cli.Command{
	Name:  "feed",
	Usage: "view your local posts",
	Action: func(c *cli.Context) error {
		dir := storage.ResolvePath(StorageDir)
		postPath := filepath.Join(dir, "posts")
		files, err := ioutil.ReadDir(postPath)
		if err != nil {
			return fmt.Errorf("no posts found or can't read posts directory: %w", err)
		}
		for _, file := range files {
			data, err := ioutil.ReadFile(filepath.Join(postPath, file.Name()))
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to read post %s: %v\n", file.Name(), err)
				continue
			}
			var p core.Post
			err = json.Unmarshal(data, &p)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid post format in %s: %v\n", file.Name(), err)
				continue
			}
			fmt.Printf("[%s] %s\n\n", time.Unix(p.Timestamp, 0).Format(time.RFC822), p.Content)
		}
		return nil
	},
}
