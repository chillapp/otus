package usecase

import (
	"otus/modules/common/health"
)

type Health struct {
}

func NewHeathUC() *Health {
	return &Health{}
}

func (uc *Health) Health() health.Health {
	return health.Health{Status: "OK"}
}
