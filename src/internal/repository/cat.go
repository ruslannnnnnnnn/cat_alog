package repository

import "cat_alog/src/internal/model"

type CatRepositoryInterface interface {
	Insert(cat *model.Cat) (err error)
	GetById(id string) (model.Cat, error)
	GetAllCats(page uint64, perPage uint32) ([]model.Cat, error)
}
