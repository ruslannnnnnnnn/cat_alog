package cassandra

import (
	"cat_alog/internal/domain/model"
	"fmt"
	"github.com/gocql/gocql"
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

	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		return cat, fmt.Errorf("invalid UUID: %v", err)
	}

	session, err := GetCassandraSession()
	if err != nil {
		return cat, err
	}
	defer session.Close()

	query := session.Query(
		"SELECT CAST(id AS TEXT), name, date_of_birth, image_url FROM catalog.cats WHERE id=?",
		uuid,
	)

	err = query.Scan(
		&cat.Id,
		&cat.Name,
		&cat.DateOfBirth,
		&cat.ImageUrl,
	)

	if err != nil {
		if err == gocql.ErrNotFound {
			return cat, fmt.Errorf("cat not found")
		}
		return cat, fmt.Errorf("scan failed: %v", err)
	}

	return cat, nil
}
