package main

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/app"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	appConfig := utils.AppConfig
	if appConfig.Debug() {
		log.SetLevel(log.DebugLevel)
	}
	application := app.CreateApp()
	server, err := tcpserver.NewServer(appConfig.ListenAddress())
	if err != nil {
		log.Error("Server failed to start")
	}
	server.SetRequestHandler(application.BinaryDataHandler())
	err = server.Listen()
	if err != nil {
		log.Error("Server failed to start listening")
	}
	log.Info("Server successfully started")
	err = server.Serve()
	if err != nil {
		log.Error("Server failed to serve")
	}
}
