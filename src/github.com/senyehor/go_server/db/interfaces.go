package db

import "github.com/senyehor/go_server/data_models"

type packet interface {
	Values() *data_models.PacketValues
	TimeInterval() int
	PacketNum() int
	DeviceID() int
}

type queryResult interface {
	RowsAffected() int64
	String() string
	Insert() bool
	Update() bool
	Delete() bool
	Select() bool
}
