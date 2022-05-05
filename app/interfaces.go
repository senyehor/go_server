package app

import "context"

type dbConnection interface {
	ExecuteWithNoReturn(context context.Context, query string) error
}

type QueryResult interface {
	RowsAffected() int64
	String() string
	Insert() bool
	Update() bool
	Delete() bool
	Select() bool
}
