package storage

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Storage interface {
	Begin() (*sqlx.Tx, error)
	Get(target interface{}, sql string, params ...interface{}) error
	Select(target interface{}, sql string, params ...interface{}) error
	LastInsertId(target *int, r sql.Result) error
}
