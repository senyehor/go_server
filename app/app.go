package app

import (
	"context"
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/data_models"
	"github.com/senyehor/go_server/db"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"net"
)

type App struct {
	connection dbConnection
}

func CreateApp() *App {
	return &App{connection: db.GetDB(utils.DBConfig)}
}

func (a *App) BinaryDataHandler() func(conn tcpserver.Connection) {
	handler := func(conn tcpserver.Connection) {
		parsedPacket, err := tryParsePacketFromIncomingData(conn)
		if err != nil {
			log.Errorf("failed to parse packet from %v", conn.GetServerAddr().IP)
			return
		}
		err = a.savePacket(parsedPacket)
		if err != nil {
			log.Errorf("failed to save packet from %v", conn.GetServerAddr().IP)
			return
		}
		a.confirmPacketProcessed(conn)
		_ = conn.Close()
	}
	return handler
}

func (a *App) savePacket(packet *data_models.Packet) error {
	queryStringToInsertPacket := composeQueryToInsertPacket(packet)
	err := a.connection.ExecuteWithNoReturn(context.Background(), queryStringToInsertPacket)
	if err != nil {
		log.Debug(err)
		return err
	}
	log.Debug("packet was inserted into database")
	return nil
}

func (a *App) confirmPacketProcessed(conn net.Conn) {
	_, err := conn.Write(composeConfirmationMessage())
	if err != nil {
		log.Error(err)
		log.Error("failed to send confirmation")
	}
	log.Debugf("confirmed packet was successfully processed from %v", conn.RemoteAddr().String())
}
