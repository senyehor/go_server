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

var (
	dbPool       = db.Connect()
	packetConfig = utils.GetPacketConfig()
)

func main() {
	config := utils.GetAppConfig()
	server, err := tcpserver.NewServer("0.0.0.0:" + config.Port())
	defer dbPool.Close()
	if err != nil {
		log.Error("Server failed to start")
	}
	server.SetRequestHandler(handleConnection)
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

func handleConnection(incomingConnection tcpserver.Connection) {
	// connection is automatically closed by framework
	// add get_packet
	log.Info("I received something")
	data, err := bufio.NewReader(incomingConnection).ReadBytes(byte(packetConfig.DataTerminator()))
	if err != nil {
		incomingConnection.Close()
	}
	packet := parser.NewPacket(data)
	// todo add packet, err
	if packet == nil {
		return
	}
	log.Info("I successfully parsed packet")
	// add function save_packet
	tmp := db.ComposeQueryString(packet)
	_, err = dbPool.Exec(context.Background(), tmp)
	if err != nil {
		return
	}
	log.Info("I wrote to db")
	// add confirm_packet_processed
	_, err = incomingConnection.Write([]byte(packetConfig.Response()))
	if err != nil {
		log.Error("failed to send answer back")
	}
	log.Info("Successfully processed data")
}
