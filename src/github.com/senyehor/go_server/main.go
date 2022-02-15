package main

import (
	"bufio"
	"context"
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/db"
	"github.com/senyehor/go_server/parser"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
)

var dbPool = db.Connect()

func main() {
	config := utils.GetServerConfig()
	server, err := tcpserver.NewServer("127.0.0.1:" + config.Port())
	defer dbPool.Close(context.Background())
	if err != nil {
		log.Error("Server failed to start")
	}
	server.SetRequestHandler(processData)
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

func processData(incomingConnection tcpserver.Connection) {
	packetConfig := utils.GetPacketConfig()
	data, err := bufio.NewReader(incomingConnection).ReadBytes(byte(packetConfig.DataTerminator()))
	if err != nil {
		incomingConnection.Close()
	}
	packet := parser.NewPacket(data)
	if packet == nil {
		incomingConnection.Close()
		return
	}
	tmp := db.ComposeQueryString(packet)
	_, err = dbPool.Exec(context.Background(), tmp)
	if err != nil {
		incomingConnection.Close()
	}
	_, err = incomingConnection.Write([]byte(packetConfig.Response()))
	if err != nil {
		log.Error("failed to send answer back")
	}
	return
}
