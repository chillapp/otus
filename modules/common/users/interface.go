package users

type Storage interface {
	Get() ([]User, error)
	Create(*User) error
}

type UseCase interface {
	Get() (interface{}, error)
	Create(*User) error
}
