package v1

import (
	"otus/modules/common/users"
	"otus/storage"
)

type Storage struct {
	db storage.Storage
}

func NewUsersStorage(db storage.Storage) *Storage {
	return &Storage{
		db: db,
	}
}

func (store *Storage) Get() []users.User {
	return nil
}

func (store *Storage) Create(user *users.User) error {
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := "INSERT INTO users SET NAME = ?"
	r, err := tx.Exec(query, user.Name)
	if err != nil {
		return err
	}
	err = store.db.LastInsertId(&user.Id, r)
	if err != nil {
		return err
	}
	return tx.Commit()
}
