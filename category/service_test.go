package category

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockCategoryRepo struct{}

var (
	Create func(category *Category) error
	List   func() ([]Category, error)
)

func (m *mockCategoryRepo) Create(category *Category) error {
	return Create(category)
}
func (m *mockCategoryRepo) List() ([]Category, error) {
	return List()
}

var categoryAppMock CategoryService = &mockCategoryRepo{}

func Test_CreateCategory(t *testing.T) {
	Create = func(category *Category) error {
		return nil
	}
	var category = &Category{
		Type: "test",
	}

	err := categoryAppMock.Create(category)
	assert.Nil(t, err)
}

func Test_ListCategory(t *testing.T) {
	List = func() ([]Category, error) {
		return []Category{}, nil
	}

	_, err := categoryAppMock.List()
	assert.Nil(t, err)
}
