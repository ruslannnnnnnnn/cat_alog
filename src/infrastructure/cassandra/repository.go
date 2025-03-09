package cassandra

import (
	"cat_alog/src/internal/model"
)

type CatRepository struct{}

func NewCatRepository() *CatRepository {
	return &CatRepository{}
}

func (c CatRepository) Insert(cat *model.Cat) (err error) {
	session, err := GetCassandraSession()
	if err != nil {
		return err
	}
	defer session.Close()
	query := session.Query("INSERT INTO catalog.cats (id, name, date_of_birth, image_url) "+
		"VALUES (?, ?, ?, ?)", cat.Id, cat.Name, cat.DateOfBirth, cat.ImageUrl,
	)
	err = query.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (c CatRepository) GetById(id string) (model.Cat, error) {
	var cat model.Cat
	session, err := GetCassandraSession()
	if err != nil {
		return cat, err
	}
	defer session.Close()
	query := session.Query("SELECT id, name, date_of_birth, image_url FROM catalog.cats WHERE id=?", id)
	err = query.Scan(cat)
	if err != nil {
		return cat, err
	}
	return cat, nil
}

func (c CatRepository) GetAllCats(page uint64, perPage uint32) ([]model.Cat, error) {
	var cats []model.Cat
	session, err := GetCassandraSession()
	if err != nil {
		return cats, err
	}
	defer session.Close()
	offset := (page - 1) * uint64(perPage)
	query := session.Query("SELECT id, name, date_of_birth, image_url FROM catalog.cats LIMIT ?, OFFSET ", perPage, offset)
	err = query.Scan(cats)
	if err != nil {
		return cats, err
	}
	return cats, nil
}
