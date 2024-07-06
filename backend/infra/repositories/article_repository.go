package repositories

import (
	"duydev.io.vn/rao-vat/app/domains"
	"duydev.io.vn/rao-vat/infra/database"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

func toEntity(data map[string]interface{}) (*domains.Article, error) {
	var entity domains.Article

	err := mapstructure.Decode(data, &entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func GetAll() ([]domains.Article, error) {
	ctx, db := database.GetDBInstance()

	iter := db.Collection("articles").Documents(ctx)

	res := []domains.Article{}

	for {
		docSnapshot, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		entity, err := toEntity(docSnapshot.Data())
		if err != nil {
			return nil, err
		}

		entity.Id = docSnapshot.Ref.ID

		res = append(res, *entity)
	}

	return res, nil
}

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
