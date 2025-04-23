package p2p

import(
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func PublishPost(data []byte) error{
	h, err := libp2p.New()
	if err != nil{
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

	if err := topic.Publish(ctx, data); err != nil{
		return err
	}

	fmt.Println("[*] Broadcast complete.")
	return nil
}