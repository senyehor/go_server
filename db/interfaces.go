package db

type DBConfig interface {
	Username() string
	Password() string
	Host() string
	Port() string
	Name() string
}
