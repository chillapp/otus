package users

import (
	"errors"
)

var ErrUserName = errors.New("username cannot be empty")
var ErrUserId = errors.New("user id cannot be empty")
var ErrUserNotFound = errors.New("user id = %d not found")
var ErrUserIsGod = errors.New("user id = 1 is god")
var ErrModifyDataEmpty = errors.New("modify data cannot be empty")
