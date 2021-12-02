package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"otus/storage"
)

type MySql struct {
	db *sqlx.DB
}

// InitStore
// Соединение с БД
func InitStore() (*MySql, error) {
	storeURI := os.Getenv("STORE_URI")
	if storeURI == "" {
		return nil, storage.ErrStoreUriEmpty
	}
	storeDB := os.Getenv("STORE_DB")
	if storeDB == "" {
		return nil, storage.ErrStoreDbEmpty
	}
	log.Println("Opening store: ", storeURI)
	db, err := sqlx.Open("mysql", storeURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dbQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", storeDB)
	_, err = db.Exec(dbQuery)
	if err != nil {
		return nil, err
	}
	err = db.Close()
	if err != nil {
		return nil, err
	}
	db, err = sqlx.Open("mysql", storeURI+storeDB)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Store opened...")
	store := &MySql{
		db: db,
	}
	return store, store.doMigrate()
}

func (store *MySql) doMigrate() error {
	_, err := store.db.Exec("CREATE TABLE IF NOT EXISTS users (id int auto_increment primary key, name varchar(256))")
	return err
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

func (store *MySql) Select(target interface{}, sql string, params ...interface{}) error {
	return store.db.Select(target, sql, params...)
}

func (store *MySql) LastInsertId(target *int, r sql.Result) error {
	lastId, err := r.LastInsertId()
	if err != nil {
		return err
	}
	*target = int(lastId)
	return nil
}
