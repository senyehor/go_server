package app

import (
	"bufio"
	"github.com/maurice2k/tcpserver"
	log "github.com/sirupsen/logrus"
)

func getBinaryDataFromConnection(incomingConnection *tcpserver.Connection) ([]byte, error) {
	data, err := bufio.NewReader(*incomingConnection).ReadBytes(byte(packetConfig.DataTerminator()))
	if err != nil {
		return nil, err
	}
	log.Debug("I received some data from connection")
	return data, nil
}

func confirmPacketProcessed(incomingConnection *tcpserver.Connection) {
	_, err := (*incomingConnection).Write([]byte(packetConfig.Response()))
	if err != nil {
		log.Error("failed to send confirmation")
	}
	log.Info("Confirmed packet was processed successfully")
}
