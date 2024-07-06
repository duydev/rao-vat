package http

import (
	"duydev.io.vn/rao-vat/infra/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
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
