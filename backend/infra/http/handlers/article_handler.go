package handlers

import (
	"net/http"

	"duydev.io.vn/rao-vat/app/domains"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (h ArticleHandler) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (h ArticleHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (h ArticleHandler) Create(ctx *gin.Context) {
	var article domains.Article

	err := ctx.BindJSON(&article)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (h ArticleHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (h ArticleHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}
