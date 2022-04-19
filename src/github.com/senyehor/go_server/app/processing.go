package app

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/db"
	"github.com/senyehor/go_server/packet"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
)

var (
	packetConfig = utils.GetPacketConfig()
)

func ProcessIncomingPacket(incomingConnection tcpserver.Connection) {
	parsedPacket, err := tryParsePacketFromIncomingData(&incomingConnection)
	if err != nil {
		log.Debug("failed to parse packet")
		return
	}
	err = db.SavePacket(parsedPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return
	}
	confirmPacketProcessed(&incomingConnection)
}

func tryParsePacketFromIncomingData(incomingConnection *tcpserver.Connection) (*packet.Packet, error) {
	rawData, err := getBinaryDataFromConnection(incomingConnection)
	if err != nil {
		return nil, err
	}
	result, err := packet.ParseFromBinary(rawData)
	if err != nil {
		return nil, err
	}
	return result, nil
}
