package repository

import (
	"cat_alog/internal/domain/model"
)

type CatRepositoryInterface interface {
	Insert(cat *model.Cat) error
	GetById(id string) (model.Cat, error)
	Search(text string) ([]model.Cat, error)
}
