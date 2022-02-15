package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/senyehor/go_server/parser"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func Connect() *pgx.Conn {
	connection, err := pgx.Connect(context.Background(), getConnString())
	if err != nil {
		log.Error(err)
		log.Error("Could not connect no db")
		os.Exit(1)
	}
	err = connection.Ping(context.Background())
	if err != nil {
		log.Error("Failed to ping db")
		os.Exit(1)
	}
	return connection
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
