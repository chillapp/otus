package v1

import (
	"otus/modules/common/users"
)

type Users struct {
	storage users.Storage
}

func NewUsersUC(storage users.Storage) *Users {
	return &Users{
		storage: storage,
	}
}

func (uc *Users) Get() interface{} {
	return uc.storage.Get()
}

func (uc *Users) Create(user *users.User) error {
	if user.Name == "" {
		return users.ErrUserName
	}
	return uc.storage.Create(user)
}
