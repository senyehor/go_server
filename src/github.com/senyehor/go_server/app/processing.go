package app

import (
	"bufio"
	"context"
	"github.com/maurice2k/tcpserver"
	database "github.com/senyehor/go_server/db"
	"github.com/senyehor/go_server/packet_parser"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
)

var (
	dbConnection = database.GetConnection()
	packetConfig = utils.GetPacketConfig()
)

func ProcessIncomingPacket(incomingConnection tcpserver.Connection) {
	packet, err := tryParsePacketFromIncomingData(&incomingConnection)
	if err != nil {
		log.Debug("failed to parse packet")
		return
	}
	err = savePacket(packet)
	if err != nil {
		log.Debug("failed to save packet")
		return
	}
	confirmPacketProcessed(&incomingConnection)
}

func tryParsePacketFromIncomingData(incomingConnection *tcpserver.Connection) (*packet_parser.Packet, error) {
	rawData, err := getBinaryDataFromConnection(incomingConnection)
	if err != nil {
		return nil, err
	}
	packet, err := packet_parser.ParseFromBinary(rawData)
	if err != nil {
		return nil, err
	}
	return packet, nil
}

func getBinaryDataFromConnection(incomingConnection *tcpserver.Connection) ([]byte, error) {
	data, err := bufio.NewReader(*incomingConnection).ReadBytes(byte(packetConfig.DataTerminator()))
	if err != nil {
		return nil, err
	}
	log.Debug("I received some data from connection")
	return data, nil
}

func savePacket(packet *packet_parser.Packet) error {
	log.Error("Db address --------------------")
	log.Error(&dbConnection)
	queryStringToInsertPacket := database.ComposeQueryStringToInsertPacket(packet)
	_, err := dbConnection.Exec(context.Background(), queryStringToInsertPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return err
	}
	log.Debug("Packet was inserted into db")
	return nil
}

func confirmPacketProcessed(incomingConnection *tcpserver.Connection) {
	_, err := (*incomingConnection).Write([]byte(packetConfig.Response()))
	if err != nil {
		log.Error("failed to send confirmation")
	}
	log.Info("Confirmed packet was processed successfully")
}
