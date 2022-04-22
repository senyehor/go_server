package db

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var database = getConnection()

func SavePacket(packet ipacket) error {
	queryStringToInsertPacket := composeQueryStringToInsertPacket(packet)
	_, err := database.Exec(context.Background(), queryStringToInsertPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return err
	}
	log.Debug("packet was inserted into db")
	return nil
}

func composeQueryStringToInsertPacket(packetToInsert ipacket) string {
	insertPart := "insert into sensor_values" +
		" (sensor_value, value_accumulation_period, package_number, boxes_set_id)"
	valuesPart := " values "
	iterator := packetToInsert.Values().Iterator()
	for iterator.HasNext() {
		valuesPart += fmt.Sprintf(
			"(%v, %v, %v, "+
				"(select boxes_set_id from boxes_sets bs join boxes b "+
				"on bs.box_id=b.box_id and box_number='%v' and bs.sensor_number=%v))",
			iterator.Value(), packetToInsert.TimeInterval(), packetToInsert.PacketNum(),
			packetToInsert.DeviceID(), iterator.ValuePosition()+1)
		if iterator.IsLast() {
			valuesPart += ";"
		} else {
			valuesPart += ", "
		}
	}
	return insertPart + valuesPart
}
