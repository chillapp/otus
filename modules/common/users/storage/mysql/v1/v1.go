package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

func (store *Storage) GetUsers() ([]users.User, error) {
	var u []users.User
	query := "SELECT id, name FROM users"
	err := store.db.Select(&u, query)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (store *Storage) GetUser(userId int) (users.User, error) {
	var u users.User
	query := "SELECT id, name FROM users WHERE id = ?"
	err := store.db.Get(&u, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			notFound := fmt.Sprintf(users.ErrUserNotFound.Error(), userId)
			return u, errors.New(notFound)
		}
		return u, err
	}
	return u, nil
}

func (store *Storage) CreateUser(user *users.User) error {
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

func (store *Storage) DeleteUser(userId int) error {
	_, err := store.GetUser(userId)
	if err != nil {
		return err
	}
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := "DELETE FROM users WHERE ID = ?"
	_, err = tx.Exec(query, userId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (store *Storage) ModifyUser(userId int, modifyData map[string]string) error {
	_, err := store.GetUser(userId)
	if err != nil {
		return err
	}
	userParams := []string{"name"}
	var queryParams []string
	var queryValues []interface{}
	for _, param := range userParams {
		if modifyData[param] != "" {
			queryParams = append(queryParams, "name = ?")
			queryValues = append(queryValues, modifyData[param])
		}
	}
	if len(queryParams) == 0 {
		return users.ErrModifyDataEmpty
	}
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			%s
		WHERE id = ?
	`, strings.Join(queryParams, ",\n"))
	queryValues = append(queryValues, userId)
	_, err = tx.Exec(query, queryValues...)
	if err != nil {
		return err
	}
	return tx.Commit()
}
