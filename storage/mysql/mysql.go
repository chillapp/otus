package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySql struct {
	db *sqlx.DB
}

func InitStore() (*MySql, error) {
	cfg := map[string]string{
		"db":    os.Getenv("STORE_DB"),
		"pwd":   os.Getenv("STORE_PWD"),
		"host":  os.Getenv("STORE_HOST"),
		"port":  os.Getenv("STORE_PORT"),
		"login": os.Getenv("STORE_LOGIN"),
	}
	dbURI := createURI(cfg)
	db, err := sqlx.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}
	return &MySql{
		db: db,
	}, nil
}

func createURI(cfg map[string]string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg["login"], cfg["pwd"], cfg["host"], cfg["port"], cfg["db"])
}

func (store *MySql) Begin() (*sqlx.Tx, error) {
	return store.db.Beginx()
}

func (store *MySql) Get(target interface{}, sql string, params ...interface{}) error {
	return store.db.Get(target, sql, params...)
}

func (store *MySql) LastInsertId(target *int, r sql.Result) error {
	lastId, err := r.LastInsertId()
	if err != nil {
		return err
	}
	*target = int(lastId)
	return nil
}
