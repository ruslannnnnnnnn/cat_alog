package service

import (
	"cat_alog/internal/domain/model"
	"cat_alog/internal/domain/repository"
	"github.com/google/uuid"
)

type CatService struct {
	repository repository.CatRepositoryInterface
}

func NewCatService(repository repository.CatRepositoryInterface) CatService {
	return CatService{repository: repository}
}

func (c CatService) GetById(id string) (model.Cat, error) {
	valid := uuid.Validate(id)
	if valid != nil {
		return model.Cat{}, valid
	}
	result, err := c.repository.GetById(id)
	if err != nil {
		return model.Cat{}, err
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

func (c CatService) Search(text string) ([]model.Cat, error) {
	result, err := c.repository.Search(text)
	if err != nil {
		return []model.Cat{}, err
	}
	return result, nil
}
