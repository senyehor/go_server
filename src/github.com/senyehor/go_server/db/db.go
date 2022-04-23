package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/senyehor/go_server/app"
)

type DB struct {
	conn *pgxpool.Pool
}

func GetDB() *DB {
	return &DB{conn: getConnection()}
}

func (db *DB) Execute(query string) (app.QueryResult, error) {
	// todo possible timeout
	return db.conn.Exec(context.Background(), query)
}
