package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h HealthCheckHandler) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
