package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/senyehor/go_server/parser"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func Connect() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(getConnString())
	if err != nil {
		log.Error(err)
		log.Error("Could not parse config")
		os.Exit(1)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Error(err)
		log.Error("Could connect :(")
		os.Exit(1)
	}
	return pool
}

func getConnString() string {
	dbConfig := utils.GetDBConfig()
	return "postgres://" +
		dbConfig.Username() + ":" +
		dbConfig.Password() + "@" +
		dbConfig.Host() + ":" +
		dbConfig.Port() +
		"/" + dbConfig.Name()
}

func ComposeQueryString(packet *parser.Packet) string {
	insertPart := "insert into sensor_values" +
		" ( sensor_value, value_accumulation_period, package_number, boxes_set_id) values "
	valuesPart := ""
	for i := 0; i < parser.PacketValuesCount; i++ {
		valuesPart += fmt.Sprintf("(%.1f, %v, %v, (select boxes_set_id from"+
			" boxes_sets bs join boxes b "+
			"on bs.box_id=b.box_id "+
			"and box_number='%v' and bs.sensor_number=%v))",
			packet.Values()[i], packet.Time(), packet.PacketNum(), packet.DeviceID(), i+1)
		if i == parser.PacketValuesCount-1 {
			valuesPart += ";"
		} else {
			valuesPart += ", "
		}
	}
	return insertPart + valuesPart
}
