package app

import "context"

type dbConnection interface {
	ExecuteWithNoReturn(context context.Context, query string) error
}
