package server_controlling

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

type CommandListener struct {
	messageReceiver *redis.PubSub
}

type Command struct {
	content string
}

func NewCommand(content string) *Command {
	if content != utils.ServerControllingConfig.StopListeningCommand() ||
		content != utils.ServerControllingConfig.ResumeListeningCommand() {
		panic("wrong command received")
	}
	return &Command{content: content}
}

func (c *Command) IsResumeListening() bool {
	return c.content == utils.ServerControllingConfig.ResumeListeningCommand()
}

func (c *Command) IsStopListening() bool {
	return c.content == utils.ServerControllingConfig.StopListeningCommand()
}

func NewCommandListener(r *redis.Client) *CommandListener {
	sub := r.Subscribe(context.Background(), utils.ServerControllingConfig.ChannelName())
	iface, err := sub.Receive(context.Background())
	if err != nil {
		log.Error(err)
		panic("failed to receive subscription message")
	}
	switch iface.(type) {
	case *redis.Subscription:
		// consume successful subscription message
		break
	case *redis.Message:
	case *redis.Pong:
	default:
		panic("subscription message was expected")
	}
	return &CommandListener{
		messageReceiver: sub,
	}
}

func (cl CommandListener) TryGetCommand() *Command {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	msg, err := cl.messageReceiver.ReceiveMessage(ctx)
	// todo check behaviour
	if msg == nil {
		return nil
	}
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return NewCommand(msg.Payload)
}
