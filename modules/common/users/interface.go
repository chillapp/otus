package users

type Storage interface {
	GetUsers() ([]User, error)
	GetUser(int) (User, error)
	CreateUser(*User) error
	DeleteUser(int) error
	ModifyUser(int, map[string]string) error
}

type UseCase interface {
	GetUser(int) (interface{}, error)
	GetUsers() (interface{}, error)
	CreateUser(*User) error
	DeleteUser(int) error
	ModifyUser(int, map[string]string) error
}
