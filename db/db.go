package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	conn *pgxpool.Pool
}

func GetDB() *DB {
	return &DB{conn: getConnection()}
}

func (db *DB) ExecuteWithNoReturn(query string) error {
	// todo possible timeout
	_, err := db.conn.Exec(context.Background(), query)
	return err
}
