package main

import (
	"duydev.io.vn/rao-vat/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	healthCheckHandler := handlers.NewHealthCheckHandler()

	apiGroup := r.Group("/api")
	apiGroup.GET("/health-check", healthCheckHandler.GetAll)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
