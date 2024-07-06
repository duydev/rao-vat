package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	r.Run()
}
