// cli/commands/post.go
package commands

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	cli "github.com/urfave/cli/v2"
	"soulblog/core"
	"soulblog/crypto"
	"soulblog/storage"
)

var Post = &cli.Command{
	Name:  "post",
	Usage: "create a new blog post",
	Action: func(c *cli.Context) error {
		if c.Args().Len() == 0 {
			return fmt.Errorf("please provide post content")
		}
		dir := storage.ResolvePath(StorageDir)
		kp, err := crypto.LoadKeys(dir)
		if err != nil {
			return fmt.Errorf("failed to load keys: %w", err)
		}
		content := c.Args().Get(0)
		post := core.Post{
			AuthorID:  fmt.Sprintf("%x", kp.PublicKey[:8]),
			Content:   content,
			Timestamp: time.Now().Unix(),
		}
		msg, err := json.Marshal(post)
		if err != nil {
			return err
		}
		post.Sig = ed25519.Sign(kp.PrivateKey, msg)
		postBytes, err := json.MarshalIndent(post, "", "  ")
		if err != nil {
			return err
		}
		postPath := filepath.Join(dir, "posts")
		_ = storage.EnsureDir(postPath)
		fname := fmt.Sprintf("%d.json", post.Timestamp)
		err = os.WriteFile(filepath.Join(postPath, fname), postBytes, 0600)
		if err != nil {
			return err
		}
		fmt.Println("[+] Post created")
		return nil
	},
}
