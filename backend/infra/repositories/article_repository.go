package repositories

import (
	"duydev.io.vn/rao-vat/app/domains"
	"duydev.io.vn/rao-vat/infra/database"
	"github.com/mitchellh/mapstructure"
)

func GetAll() {}

func Get() {}

func Create(article domains.Article) (*domains.Article, error) {
	ctx, db := database.GetDBInstance()

	docRef, _, err := db.Collection("articles").Add(ctx, article)
	if err != nil {
		return nil, err
	}

	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	var retrievedArticle domains.Article

	err = mapstructure.Decode(docSnapshot.Data(), &retrievedArticle)
	if err != nil {
		return nil, err
	}

	retrievedArticle.Id = docRef.ID

	return &retrievedArticle, nil
}

func Update() {}

func Delete() {}
