// p2p/sync.go
package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"soulblog/core"
	"soulblog/storage"
)

const FeedDir = "feed"
const TopicName = "soulblog-posts"

func StartSync(baseDir string) error {
	h, err := libp2p.New()
	if err != nil {
		return err
	}
	ctx := context.Background()
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return err
	}
	topic, err := ps.Join(TopicName)
	if err != nil {
		return err
	}
	sub, err := topic.Subscribe()
	if err != nil {
		return err
	}

	feedPath := filepath.Join(baseDir, FeedDir)

	if err := storage.EnsureDir(feedPath); err != nil {
		return fmt.Errorf("failed to create feed dir: %w", err)
	}	
	
	fmt.Println("[+] Sync started. Listening for posts... (press Ctrl+C to stop)")

	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			fmt.Println("sub error:", err)
			continue
		}
		var p core.Post
		err = json.Unmarshal(msg.Data, &p)
		if err != nil {
			fmt.Println("invalid post format")
			continue
		}
		fname := fmt.Sprintf("%d.json", p.Timestamp)
		postFile := filepath.Join(feedPath, fname)
		os.WriteFile(postFile, msg.Data, 0600)
		fmt.Printf("[+] Received post from %s\n", p.AuthorID)
	}
}
