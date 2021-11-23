package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"otus/modules/common/users"
	"otus/tools/http/response"
)

type Handler struct {
	users users.UseCase
}

func NewHandler(users users.UseCase) *Handler {
	return &Handler{
		users: users,
	}
}

func (h *Handler) Get(ctx *gin.Context) {
	data := h.users.Get()
	response.WithData(ctx, data)
}

func (h *Handler) Create(ctx *gin.Context) {
	user := new(users.User)
	err := ctx.BindJSON(user)
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.users.Create(user)
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.WithData(ctx, user)
}
