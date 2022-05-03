package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	conn *pgxpool.Pool
}

func GetDB(config DBConfig) *DB {
	return &DB{conn: getConnection(composeConnectionString(config))}
}

func (db *DB) ExecuteWithNoReturn(context context.Context, query string) error {
	_, err := db.conn.Exec(context, query)
	return err
}
