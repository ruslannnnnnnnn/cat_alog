package service

import (
	"cat_alog/internal/domain/model"
	testee "cat_alog/internal/domain/service"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	validImageURL   = "https://upload.wikimedia.org/wikipedia/commons/4/4d/Cat_November_2010-1a.jpg"
	validUUID       = "c36dde92-1242-11f0-b7fc-0242ac120003"
	invalidUUID     = "invalid-uuid"
	nonExistentUUID = "c36dde93-1242-11f0-b7fc-0242ac120003"
)

var testCats = []model.Cat{
	{Id: validUUID, Name: "pushok", ImageUrl: validImageURL},
	{Id: "c36dde92-1242-11f0-b7fc-0242ac120004", Name: "Barsik"},
	{Id: "c36dde92-1242-11f0-b7fc-0242ac120005", Name: "Asteroid destroyer"},
}

type MockCatRepository struct {
	mock.Mock
}

func (m *MockCatRepository) Search(query string) ([]model.Cat, error) {
	args := m.Called(query)
	return args.Get(0).([]model.Cat), args.Error(1)
}

func (m *MockCatRepository) GetById(id string) (model.Cat, error) {
	args := m.Called(id)
	return args.Get(0).(model.Cat), args.Error(1)
}

func (m *MockCatRepository) Insert(cat *model.Cat) error {
	args := m.Called(cat)
	return args.Error(0)
}

func newTestService(t *testing.T) (testee.CatService, *MockCatRepository) {
	t.Helper()
	repo := &MockCatRepository{}
	return testee.NewCatService(repo), repo
}

func TestCreateCat(t *testing.T) {
	service, repo := newTestService(t)

	t.Run("valid cat", func(t *testing.T) {
		cat := model.Cat{
			Id:       validUUID,
			Name:     "kotik",
			ImageUrl: validImageURL,
		}

		repo.On("Insert", &cat).Return(nil).Once()

		err := service.Create(&cat)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid data", func(t *testing.T) {
		invalidCat := model.Cat{
			Id:       invalidUUID,
			Name:     "",
			ImageUrl: "invalid_url",
		}

		err := service.Create(&invalidCat)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cat has invalid image url")
	})
}

func TestGetCatById(t *testing.T) {
	service, repo := newTestService(t)

	t.Run("existing cat", func(t *testing.T) {
		expectedCat := testCats[0]
		repo.On("GetById", validUUID).Return(expectedCat, nil).Once()

		cat, err := service.GetById(validUUID)
		require.NoError(t, err)
		assert.Equal(t, expectedCat, cat)
	})

	t.Run("nonexistent cat", func(t *testing.T) {
		repo.On("GetById", nonExistentUUID).Return(model.Cat{}, fmt.Errorf("not found")).Once()

		_, err := service.GetById(nonExistentUUID)
		assert.Error(t, err)
	})

	t.Run("invalid UUID", func(t *testing.T) {
		_, err := service.GetById(invalidUUID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid UUID")
	})
}

func TestSearchCat(t *testing.T) {
	service, repo := newTestService(t)

	testCases := []struct {
		name        string
		query       string
		mockResults []model.Cat
		mockError   error
		expected    []model.Cat
		expectError bool
	}{
		{
			name:        "successful search",
			query:       "pushok",
			mockResults: testCats[:1],
			expected:    testCats[:1],
		},
		{
			name:        "partial match",
			query:       "destroyer",
			mockResults: testCats[2:],
			expected:    testCats[2:],
		},
		{
			name:        "empty results",
			query:       "unknown",
			mockResults: []model.Cat{},
			expected:    []model.Cat{},
		},
		{
			name:        "repository error",
			query:       "error",
			mockError:   fmt.Errorf("database error"),
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo.On("Search", tc.query).Return(tc.mockResults, tc.mockError).Once()

			result, err := service.Search(tc.query)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
