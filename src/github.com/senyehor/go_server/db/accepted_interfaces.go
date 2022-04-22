package db

import "github.com/senyehor/go_server/data_models"

type ipacket interface {
	Values() *data_models.PacketValues
	TimeInterval() int
	PacketNum() int
	DeviceID() int
}
