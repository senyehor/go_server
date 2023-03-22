package server_controlling

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

func CreateServer(f tcpserver.RequestHandlerFunc) *tcpserver.Server {
	server, err := tcpserver.NewServer(utils.AppConfig.ListenAddress())
	if err != nil {
		log.Error(err)
		panic("Server failed to start")
	}
	server.SetRequestHandler(f)
	return server
}

func RunServer(s *tcpserver.Server) {
	err := s.Listen()
	if err != nil {
		log.Error(err)
		panic("Server failed to start listening")
	}
	log.Info("Server successfully started")
	err = s.Serve()
	if err != nil {
		log.Error(err)
		panic("Server failed to serve")
	}
}

func StopServer(s *tcpserver.Server) {
	err := s.Shutdown(time.Second * 10)
	if err != nil {
		log.Error(err)
		panic("Server failed to shutdown")
	}
}
