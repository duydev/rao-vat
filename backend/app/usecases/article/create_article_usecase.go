package usecases

import (
	"duydev.io.vn/rao-vat/app/domains"
	"duydev.io.vn/rao-vat/infra/repositories"
)

func CreateArticleUsecase(article domains.Article) (*domains.Article, error) {
	createdArticle, err := repositories.Create(article)

	return createdArticle, err
}
