package app

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/binary_parser"
	"github.com/senyehor/go_server/data_models"
	"github.com/senyehor/go_server/db"
	log "github.com/sirupsen/logrus"
)

func ProcessIncomingRawPacket(incomingConnection tcpserver.Connection) {
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

func tryParsePacketFromIncomingData(incomingConnection *tcpserver.Connection) (*data_models.Packet, error) {
	rawData, err := getBinaryDataFromConnection(incomingConnection)
	if err != nil {
		return nil, err
	}
	result, err := binary_parser.ParseFromBinary(rawData)
	if err != nil {
		return nil, err
	}
	return result, nil
}
