package service

import (
	"cat_alog/src/infrastructure/cassandra"
	"cat_alog/src/internal/model"
)

type CatService struct{}

func NewCatService() *CatService {
	return &CatService{}
}

func (c CatService) GetById(id string) (model.Cat, error) {
	catRepo := cassandra.NewCatRepository()
	result, err := catRepo.GetById(id)
	if err != nil {
		return model.Cat{}, err
	}
	return result, nil
}

func (c CatService) GetAll(page uint, perPage uint) ([]model.Cat, error) {
	catRepo := cassandra.NewCatRepository()
	result, err := catRepo.GetAllCats(page, perPage)
	if err != nil {
		return []model.Cat{}, err
	}
	return result, nil
}

func (c CatService) Create(cat *model.Cat) error {
	catRepo := cassandra.NewCatRepository()
	err := catRepo.Insert(cat)
	if err != nil {
		return err
	}
	return nil
}
