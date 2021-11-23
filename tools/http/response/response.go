package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Route   string      `json:"route,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WithSuccess(ctx *gin.Context, data interface{}) {
	resp := response{}
	resp.Success = true
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
}

func WithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func WithFail(ctx *gin.Context, status int, message string) {
	resp := response{}
	resp.Success = false
	resp.Message = message
	ctx.JSON(http.StatusOK, resp)
}
