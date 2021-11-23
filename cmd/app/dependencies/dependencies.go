package dependencies

import (
	"otus/modules/common/health"
	"otus/modules/common/health/usecase"
	"otus/modules/common/users"
	usersStore "otus/modules/common/users/storage/mysql/v1"
	usersUC "otus/modules/common/users/usecase/v1"
	"otus/storage"
)

type Modules struct {
	Health health.UseCase
	Users  users.UseCase
}

func Injection(storage storage.Storage) *Modules {
	modules := &Modules{}
	modules.Health = usecase.NewHeathUC()

	usersStorage := usersStore.NewUsersStorage(storage)
	modules.Users = usersUC.NewUsersUC(usersStorage)
	return modules
}
