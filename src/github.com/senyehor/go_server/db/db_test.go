package db

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}

type DBTestSuite struct {
	suite.Suite
}

type PackedMock struct {
	mock.Mock
}

func (d *DBTestSuite) TestSavePacket() {
	tx, err := database.Begin(context.Background())
	if err != nil {
		d.Fail("could not begin transaction")
		return
	}
	defer tx.Rollback(context.Background())
	//var values []string
	rand.Seed(time.Now().UnixMilli())

}
