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
	for iterationItem := range packetToInsert.Values().Iterate() {
		valuesPart += fmt.Sprintf("(%v, %v, %v, (select boxes_set_id from"+
			" boxes_sets bs join boxes b "+
			"on bs.box_id=b.box_id "+
			"and box_number='%v' and bs.sensor_number=%v))",
			iterationItem.Value(), packetToInsert.TimeInterval(), packetToInsert.PacketNum(),
			packetToInsert.DeviceID(), iterationItem.ValuePosition()+1)
		if iterationItem.IsLast() {
			valuesPart += ";"
		} else {
			valuesPart += ", "
		}
	}
	return insertPart + valuesPart
}
