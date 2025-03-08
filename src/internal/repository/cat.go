package repository

import "cat_alog/src/internal/model"

type CatRepositoryInterface interface {
	Insert(cat *model.Cat) (err error)
	GetById(id string) (model.Cat, error)
	GetAllCats(page int, perPage int) ([]model.Cat, error)
}
