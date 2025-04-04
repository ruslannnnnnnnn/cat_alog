package service

import (
	"cat_alog/src/internal/model"
	"cat_alog/src/internal/repository"
	"errors"
)

type CatService struct {
	repository repository.CatRepositoryInterface
}

func NewCatService(repository repository.CatRepositoryInterface) CatService {
	return CatService{repository: repository}
}

func (c CatService) GetById(id string) (model.Cat, error) {
	result, err := c.repository.GetById(id)
	if err != nil {
		return model.Cat{}, err
	}
	return result, nil
}

func (c CatService) GetAll(page uint64, perPage uint32) ([]model.Cat, error) {
	if page < 1 {
		return []model.Cat{}, errors.New("page must be greater than zero")
	}
	result, err := c.repository.GetAllCats(page, perPage)
	if err != nil {
		return []model.Cat{}, err
	}
	return result, nil
}

func (c CatService) Create(cat *model.Cat) error {
	err := cat.IsValid()
	if err != nil {
		return err
	}
	err = c.repository.Insert(cat)
	if err != nil {
		return err
	}
	return nil
}
