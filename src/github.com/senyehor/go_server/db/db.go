package db

import (
	log "github.com/sirupsen/logrus"
)

// todo make db struct

func SavePacket(packet packet) error {
	queryStringToInsertPacket := composeQueryStringToInsertPacket(packet)
	_, err := executeQuery(queryStringToInsertPacket)
	if err != nil {
		log.Debug("failed to save packet")
		return err
	}
	log.Debug("packet was inserted into db")
	return nil
}
