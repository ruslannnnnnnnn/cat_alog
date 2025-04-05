package repository

import (
	"cat_alog/internal/domain/model"
)

type CatRepositoryInterface interface {
	Insert(cat *model.Cat) (err error)
	GetById(id string) (model.Cat, error)
}
