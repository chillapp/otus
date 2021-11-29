package http

import (
	"github.com/gin-gonic/gin"
	"otus/modules/common/users"
)

func RegUsersHTTP(router *gin.Engine, users users.UseCase) {
	h := NewHandler(users)
	router.GET("users", h.GetUsers)
	router.POST("user", h.CreateUser)
	router.GET("user/:id", h.GetUser)
	router.PUT("user/:id", h.ModifyUser)
	router.DELETE("user/:id", h.DeleteUser)
}
