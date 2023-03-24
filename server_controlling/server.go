package server_controlling

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	Running  = "running"
	Inactive = "inactive"
)

type Server struct {
	server  *tcpserver.Server
	status  string
	wg      sync.WaitGroup
	handler tcpserver.RequestHandlerFunc
}

func (s *Server) Run() {
	s.server = createServer(s.handler)
	err := s.server.Listen()
	if err != nil {
		log.Error(err)
		panic("Server failed to listen")
	}
	s.run()
}

func (s *Server) Stop() {
	if s.isInactive() {
		return
	}
	err := s.server.Shutdown(time.Second * 10)
	s.wg.Done()
	if err != nil {
		log.Error(err)
		panic("Server failed to shutdown")
	}
	s.setInactive()
}

func (s *Server) run() {
	s.wg.Add(1)
	s.setRunning()
	go func(s *tcpserver.Server) {
		err := s.Serve()
		if err != nil {
			log.Error(err)
			panic("Server failed to start")
		}
	}(s.server)
}

func (s *Server) isRunning() bool {
	return s.status == Running
}
func (s *Server) setRunning() {
	s.status = Running
}
func (s *Server) setInactive() {
	s.status = Inactive
}
func (s *Server) isInactive() bool {
	return s.status == Inactive
}

func createServer(f tcpserver.RequestHandlerFunc) *tcpserver.Server {
	server, err := tcpserver.NewServer(utils.AppConfig.ListenAddress())
	if err != nil {
		log.Error(err)
		panic("server creation failed")
	}
	server.SetRequestHandler(f)
	return server
}

func CreateServer(f tcpserver.RequestHandlerFunc) *Server {
	return &Server{
		server:  nil,
		status:  Inactive,
		handler: f,
	}
}
