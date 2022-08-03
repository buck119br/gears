package sql

import (
	"database/sql"
)

type Database interface {
	Init(Config) error
	Get() (*sql.DB, error)
}
