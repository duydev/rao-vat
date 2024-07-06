package main

import (
	"duydev.io.vn/rao-vat/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	healthCheckHandler := handlers.NewHealthCheckHandler()
	articleHandler := handlers.NewArticleHandler()

	apiGroup := r.Group("/api")

	apiGroup.GET("/health-check", healthCheckHandler.GetAll)

	articleGroup := apiGroup.Group("/articles")
	articleGroup.GET("/", articleHandler.GetAll)
	articleGroup.GET("/:id", articleHandler.Get)
	articleGroup.POST("/", articleHandler.Create)
	articleGroup.PUT("/:id", articleHandler.Update)
	articleGroup.DELETE("/:id", articleHandler.Delete)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
