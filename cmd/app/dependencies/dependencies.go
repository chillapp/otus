package dependencies

import (
	"otus/modules/common/health"
	"otus/modules/common/health/usecase"
)

type Modules struct {
	Health health.UseCase
}

func Injection() *Modules {
	modules := &Modules{}
	modules.Health = usecase.NewHeathUC()
	return modules
}
