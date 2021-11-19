package http

import (
	"github.com/gin-gonic/gin"
	"otus/modules/common/health"
)

func RegHealthHTTP(router *gin.Engine, useCase health.UseCase) {
	h := NewHandler(useCase)
	router.GET("", h.Hello)
	router.GET("health", h.Health)
	router.GET("version", h.Version)
}
