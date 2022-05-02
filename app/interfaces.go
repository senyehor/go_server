package app

type dbConnection interface {
	ExecuteWithNoReturn(query string) error
}

type QueryResult interface {
	RowsAffected() int64
	String() string
	Insert() bool
	Update() bool
	Delete() bool
	Select() bool
}
