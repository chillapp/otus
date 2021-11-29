package http

import (
	"net/http"
	"strconv"

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

func getIntParam(ctx *gin.Context, param string) (int, error) {
	strValue := ctx.Params.ByName(param)
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	data, err := h.users.GetUsers()
	if err != nil {
		response.WithFail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithSuccess(ctx, data)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	userId, err := getIntParam(ctx, "id")
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.users.GetUser(userId)
	if err != nil {
		response.WithFail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithSuccess(ctx, user)
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := new(users.User)
	err := ctx.BindJSON(user)
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.users.CreateUser(user)
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.WithData(ctx, user)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	userId, err := getIntParam(ctx, "id")
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.users.DeleteUser(userId)
	if err != nil {
		response.WithFail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithSuccess(ctx, nil)
}

func (h *Handler) ModifyUser(ctx *gin.Context) {
	userId, err := getIntParam(ctx, "id")
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userModifyData := map[string]string{}
	err = ctx.BindJSON(&userModifyData)
	if err != nil {
		response.WithFail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.users.ModifyUser(userId, userModifyData)
	if err != nil {
		response.WithFail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.WithSuccess(ctx, nil)
}
