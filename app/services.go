package app

import (
	"bufio"
	"fmt"
	"github.com/senyehor/go_server/binary_parser"
	"github.com/senyehor/go_server/data_models"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"net"
)

type dbConnection interface {
	ExecuteWithNoReturn(query string) error
}

type QueryResult interface {
	RowsAffected() int64
	String() string
	Insert() bool
	Update() bool
	Delete() bool
	Select() bool
}

func getBinaryDataFromConnection(incomingConnection net.Conn) ([]byte, error) {
	data, err := bufio.NewReader(incomingConnection).ReadBytes(getDataTerminator())
	if err != nil {
		return nil, err
	}
	log.Debug("I received some data from connection")
	return data, nil
}

func getDataTerminator() byte {
	return byte(utils.PacketConfig.DataTerminator())
}

func composeConfirmationMessage() []byte {
	return []byte(utils.PacketConfig.Response())
}

func tryParsePacketFromIncomingData(incomingConnection net.Conn) (*data_models.Packet, error) {
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

func composeQueryToInsertPacket(packet *data_models.Packet) string {
	insertPart := "insert into sensor_values" +
		" (sensor_value, value_accumulation_period, package_number, boxes_set_id)"
	valuesPart := " values "
	iterator := packet.Values().Iterator()
	for iterator.HasNext() {
		valuesPart += fmt.Sprintf(
			"(%v, %v, %v, "+
				"(select boxes_set_id from boxes_sets bs join boxes b "+
				"on bs.box_id=b.box_id and box_number='%v' and bs.sensor_number=%v))",
			iterator.Value(), packet.TimeInterval(), packet.PacketNum(),
			packet.DeviceID(), iterator.ValuePosition()+1)
		if iterator.IsLast() {
			valuesPart += ";"
		} else {
			valuesPart += ", "
		}
	}
	return insertPart + valuesPart
}
