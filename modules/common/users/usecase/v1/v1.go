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

func (uc *Users) GetUser(userId int) (interface{}, error) {
	if userId == 0 {
		return nil, users.ErrUserId
	}
	return uc.storage.GetUser(userId)
}

func (uc *Users) GetUsers() (interface{}, error) {
	return uc.storage.GetUsers()
}

func (uc *Users) CreateUser(user *users.User) error {
	if user.Name == "" {
		return users.ErrUserName
	}
	return uc.storage.CreateUser(user)
}

func (uc *Users) DeleteUser(userId int) error {
	if userId == 0 {
		return users.ErrUserId
	}
	if userId == 1 {
		return users.ErrUserIsGod
	}
	return uc.storage.DeleteUser(userId)
}

func (uc *Users) ModifyUser(userId int, modifyData map[string]string) error {
	if userId == 0 {
		return users.ErrUserId
	}
	if userId == 1 {
		return users.ErrUserIsGod
	}
	return uc.storage.ModifyUser(userId, modifyData)
}
