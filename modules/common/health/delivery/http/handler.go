package http

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"otus/modules/common/health"
	"otus/tools/http/response"
)

type Handler struct {
	health health.UseCase
}

func NewHandler(health health.UseCase) *Handler {
	return &Handler{
		health: health,
	}
}

func (h *Handler) Health(ctx *gin.Context) {
	status := h.health.Health()
	response.WithData(ctx, status)
}

func (h *Handler) Hello(ctx *gin.Context) {
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello"
	}
	helloMessage := fmt.Sprintf("%s world from %s!", greeting, os.Getenv("HOSTNAME"))
	response.WithData(ctx, helloMessage)
}

func (h *Handler) Version(ctx *gin.Context) {
	version := health.Version{
		Version: "0.3",
	}
	response.WithData(ctx, version)
}
