package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/senyehor/go_server/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

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
