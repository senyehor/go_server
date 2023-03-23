package server_controlling

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	Listening = "listening"
	Stopped   = "stopped"
)

type Server struct {
	server *tcpserver.Server
	status string
}

func (s *Server) Run() {
	if s.isListening() {
		return
	}
	err := s.server.Listen()
	if err != nil {
		log.Error(err)
		panic("Server failed to start listening")
	}
	log.Info("Server successfully started")
	err = s.server.Serve()
	if err != nil {
		log.Error(err)
		panic("Server failed to serve")
	}
	s.setListening()
}

func (s *Server) Stop() {
	if s.isStopped() {
		return
	}
	err := s.server.Shutdown(time.Second * 10)
	if err != nil {
		log.Error(err)
		panic("Server failed to shutdown")
	}
	s.setStopped()
}

func (s *Server) isListening() bool {
	return s.status == Listening
}
func (s *Server) setListening() {
	s.status = Listening
}
func (s *Server) setStopped() {
	s.status = Stopped
}
func (s *Server) isStopped() bool {
	return s.status == Stopped
}

func CreateServer(f tcpserver.RequestHandlerFunc) *Server {
	server, err := tcpserver.NewServer(utils.AppConfig.ListenAddress())
	if err != nil {
		log.Error(err)
		panic("Server failed to start")
	}
	server.SetRequestHandler(f)
	return &Server{
		server: server,
		status: Stopped,
	}
}
