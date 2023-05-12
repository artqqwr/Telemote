package database

type Driver string

const (
	Sqlite   Driver = "sqlite"
	Postgres Driver = "postgres"
)

type Config struct {
	Dsn    string
	Driver Driver
}
