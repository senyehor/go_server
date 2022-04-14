package db

import (
	"context"
	"fmt"
	"github.com/senyehor/go_server/packet_parser"
	log "github.com/sirupsen/logrus"
)

var database = getConnection()

func SavePacket(packet *packet_parser.Packet) error {
	queryStringToInsertPacket := composeQueryStringToInsertPacket(packet)
	_, err := database.Exec(context.Background(), queryStringToInsertPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return err
	}
	log.Debug("Packet was inserted into db")
	return nil
}

func composeQueryStringToInsertPacket(packet *packet_parser.Packet) string {
	insertPart := "insert into sensor_values" +
		" (sensor_value, value_accumulation_period, package_number, boxes_set_id)"
	valuesPart := " values "
	for index, value := range packet.Values() {
		valuesPart += fmt.Sprintf("(%.1f, %v, %v, (select boxes_set_id from"+
			" boxes_sets bs join boxes b "+
			"on bs.box_id=b.box_id "+
			"and box_number='%v' and bs.sensor_number=%v))",
			value, packet.Time(), packet.PacketNum(), packet.DeviceID(), index+1)
		if index == packet_parser.PacketValuesCount-1 { // last value in packet
			valuesPart += ";"
		} else {
			valuesPart += ", "
		}
	}
	return insertPart + valuesPart
}
