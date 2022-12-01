package sql

import (
	"database/sql"
)

type Database interface {
	Init(Config) error
	Close() error
	Get() (*sql.DB, error)
}
