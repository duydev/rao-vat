package usecases

import (
	"duydev.io.vn/rao-vat/app/domains"
	"duydev.io.vn/rao-vat/infra/repositories"
)

func GetAllArticlesUsecase() ([]domains.Article, error) {
	return repositories.GetAll()
}
