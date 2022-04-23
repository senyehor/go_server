package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

var database = getConnection()

func composeQueryStringToInsertPacket(packetToInsert packet) string {
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

func executeQuery(query string) (queryResult, error) {
	return database.Exec(context.Background(), query)
}

func getConnection() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(getConnString())
	if err != nil {
		log.Error(err)
		log.Error("Could not parse config")
		os.Exit(1)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return pool
}

func getConnString() string {
	return "postgres://" +
		utils.DBConfig.Username() + ":" +
		utils.DBConfig.Password() + "@" +
		utils.DBConfig.Host() + ":" +
		utils.DBConfig.Port() +
		"/" + utils.DBConfig.Name()
}
