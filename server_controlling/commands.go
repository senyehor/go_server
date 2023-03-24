package server_controlling

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
)

type CommandsChannel <-chan *redis.Message

type Command string

var (
	RunServer  = Command(utils.ServerControllingConfig.StartListeningCommand())
	StopServer = Command(utils.ServerControllingConfig.StopListeningCommand())
)

func NewCommand(message *redis.Message) Command {
	content := message.Payload
	switch content {
	case utils.ServerControllingConfig.StartListeningCommand():
		return RunServer
	case utils.ServerControllingConfig.StopListeningCommand():
		return StopServer
	default:
		panic("unexpected command")
	}
}
func NewCommandsChannel(r *redis.Client) CommandsChannel {
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
	return sub.Channel()
}
