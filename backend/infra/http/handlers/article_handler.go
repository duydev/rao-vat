package handlers

import (
	"net/http"

	"duydev.io.vn/rao-vat/app/domains"
	usecases "duydev.io.vn/rao-vat/app/usecases/article"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (h ArticleHandler) GetAll(ctx *gin.Context) {
	articles, err := usecases.GetAllArticlesUseCase()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": articles})
}

func (h ArticleHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (h ArticleHandler) Create(ctx *gin.Context) {
	var article domains.Article

	if err := ctx.BindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	createdArticle, err := usecases.CreateArticleUseCase(article)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": createdArticle})
}

func (h ArticleHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (h ArticleHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok"})
}
