package http

import (
	"github.com/gin-gonic/gin"
	"otus/modules/common/users"
)

func RegUsersHTTP(router *gin.Engine, users users.UseCase) {
	h := NewHandler(users)
	router.GET("users", h.Get)
	router.POST("users", h.Create)
}
