package main

import (
	"github.com/redis/go-redis/v9"
	"github.com/senyehor/go_server/app"
	"github.com/senyehor/go_server/server_controlling"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	appConfig := utils.AppConfig
	if appConfig.Debug() {
		log.SetLevel(log.DebugLevel)
	}
	application := app.CreateApp()
	server := server_controlling.CreateServer(application.BinaryDataHandler())
	server.Run()
	log.Info("server started")
	r := redis.NewClient(&redis.Options{
		Addr:     utils.RedisConfig.Address,
		Password: utils.RedisConfig.Password,
		DB:       utils.RedisConfig.DB,
	})
	commandsChannel := server_controlling.NewCommandsChannel(r)
	for {
		select {
		case msg := <-commandsChannel:
			command := server_controlling.NewCommand(msg)
			switch command {
			case server_controlling.RunServer:
				server.Run()
				log.Info("server started")
			case server_controlling.StopServer:
				server.Stop()
				log.Info("server stopped")
			}
		default:
			time.Sleep(5 * time.Second)
		}
	}
}
