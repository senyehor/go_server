package db

import (
	"context"
	"fmt"
	"github.com/senyehor/go_server/packet"
	log "github.com/sirupsen/logrus"
)

var database = getConnection()

func SavePacket(packet *packet.Packet) error {
	queryStringToInsertPacket := composeQueryStringToInsertPacket(packet)
	_, err := database.Exec(context.Background(), queryStringToInsertPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return err
	}
	log.Debug("packet was inserted into db")
	return nil
}

func composeQueryStringToInsertPacket(packetToInsert *packet.Packet) string {
	insertPart := "insert into sensor_values" +
		" (sensor_value, value_accumulation_period, package_number, boxes_set_id)"
	valuesPart := " values "
	for item := range packetToInsert.Values().Iterate() {
		valuesPart += fmt.Sprintf("(%.1f, %v, %v, (select boxes_set_id from"+
			" boxes_sets bs join boxes b "+
			"on bs.box_id=b.box_id "+
			"and box_number='%v' and bs.sensor_number=%v))",
			item.Value(), packetToInsert.Time(), packetToInsert.PacketNum(),
			packetToInsert.DeviceID(), item.ValuePosition())
		if item.IsLast() {
			valuesPart += ";"
		} else {
			valuesPart += ", "
		}
	}
	return insertPart + valuesPart
}
