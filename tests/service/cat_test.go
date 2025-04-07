package service

import (
	"cat_alog/internal/domain/model"
	testee "cat_alog/internal/domain/service"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockCatRepositoryInterface struct {
	mock.Mock
}

const firstCatId = "c36dde92-1242-11f0-b7fc-0242ac120004"

func getAllCats() ([]model.Cat, error) {
	return []model.Cat{
		{Id: firstCatId, Name: "cat"},
		{Id: "c36dde92-1242-11f0-b7fc-0242ac120003", Name: "cat"},
		{Id: "c36dde92-1242-11f0-b7fc-0242ac120006", Name: "not cat"},
	}, nil
}

func (m *MockCatRepositoryInterface) Search(text string) ([]model.Cat, error) {
	if text == "cat" {
		getAllCats()
	}
	return []model.Cat{}, nil
}

func (m *MockCatRepositoryInterface) GetById(id string) (model.Cat, error) {

	args := m.Called(id)
	if args.Get(0) == firstCatId {

	}
	return model.Cat{}, fmt.Errorf("cat not found")
}

func (m *MockCatRepositoryInterface) Insert(cat *model.Cat) (err error) {
	return nil
}

func TestCreateCat(t *testing.T) {
	repo := MockCatRepositoryInterface{}
	testeeService := testee.NewCatService(&repo)

	validCat := model.Cat{
		Id:          "c36dde92-1242-11f0-b7fc-0242ac120003",
		Name:        "test cat",
		DateOfBirth: time.Time{},
		ImageUrl:    "https://www.google.ru/url?sa=i&url=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FTabby_cat&psig=AOvVaw0LmvK3nea-64S4hifk4OFd&ust=1744137340737000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCPDotofIxowDFQAAAAAdAAAAABAE",
	}
	err := testeeService.Create(&validCat)
	if err != nil {
		t.Error(err)
	}

	invalidCat := model.Cat{
		Id:          "",
		Name:        "test cat",
		DateOfBirth: time.Time{},
		ImageUrl:    "invalid url",
	}

	err = testeeService.Create(&invalidCat)
	if err == nil {
		t.Error("Create should fail")
	}
}

func TestGetCatById(t *testing.T) {
	repo := MockCatRepositoryInterface{}
	testeeService := testee.NewCatService(&repo)

}
