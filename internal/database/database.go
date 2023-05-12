package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	config *Config
}

func New(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) Connect() *gorm.DB {
	var conn *gorm.DB
	var err error

	if db.config.Driver == Sqlite {
		conn, err = gorm.Open(sqlite.Open(db.config.Dsn))
	}
	if db.config.Driver == Postgres {
		// TODO: add postgres connection statement
		conn, err = gorm.Open(sqlite.Open(db.config.Dsn))
	}

	if err != nil {
		panic(err)
	}

	return conn
}
