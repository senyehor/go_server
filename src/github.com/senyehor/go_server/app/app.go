package app

import (
	"github.com/maurice2k/tcpserver"
	"github.com/senyehor/go_server/data_models"
	"github.com/senyehor/go_server/db"
	log "github.com/sirupsen/logrus"
	"net"
)

type App struct {
	db dbConnection
}

func CreateApp() *App {
	return &App{db: db.GetDB()}
}

func (a *App) BinaryDataHandler() func(conn tcpserver.Connection) {
	handler := func(conn tcpserver.Connection) {
		parsedPacket, err := tryParsePacketFromIncomingData(conn)
		if err != nil {
			log.Debug("failed to parse packet")
			return
		}
		err = a.savePacket(parsedPacket)
		if err != nil {
			log.Debug("failed to save packet")
			return
		}
		a.confirmPacketProcessed(conn)
	}
	return handler
}

func (a *App) savePacket(packet *data_models.Packet) error {
	queryStringToInsertPacket := composeQueryToInsertPacket(packet)
	_, err := a.db.Execute(queryStringToInsertPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return err
	}
	log.Debug("packet was inserted into db")
	return nil
}

func (a *App) confirmPacketProcessed(conn net.Conn) {
	_, err := conn.Write(composeConfirmationMessage())
	if err != nil {
		log.Error("failed to send confirmation")
	}
	log.Info("Confirmed Packet was processed successfully")
}
