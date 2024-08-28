package usecases

import (
	"duydev.io.vn/rao-vat/app/domains"
	"duydev.io.vn/rao-vat/infra/repositories"
)

func GetAllArticlesUseCase() ([]domains.Article, error) {
	return repositories.GetAll()
}
