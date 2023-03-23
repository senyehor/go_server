package main

import (
	"github.com/redis/go-redis/v9"
	"github.com/senyehor/go_server/app"
	"github.com/senyehor/go_server/server_controlling"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	appConfig := utils.AppConfig
	if appConfig.Debug() {
		log.SetLevel(log.DebugLevel)
	}
	application := app.CreateApp()
	server := server_controlling.CreateServer(application.BinaryDataHandler())
	server.Run()
	r := redis.NewClient(&redis.Options{
		Addr:     utils.RedisConfig.Address,
		Password: utils.RedisConfig.Password,
		DB:       utils.RedisConfig.DB,
	})
	log.Error("redis created")
	commandListener := server_controlling.NewCommandListener(r)
	for {
		command := commandListener.TryGetCommand()
		if command == nil {
			continue
		}
		if command.IsStopListening() {
			server.Stop()
		}
		if command.IsResumeListening() {
			server.Run()
		}
	}
}
