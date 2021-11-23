package users

type Storage interface {
	Get() []User
	Create(*User) error
}

type UseCase interface {
	Get() interface{}
	Create(*User) error
}
