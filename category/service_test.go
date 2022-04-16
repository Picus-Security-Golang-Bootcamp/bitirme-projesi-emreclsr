package category_test

import (
	"github.com/emreclsr/picusfinal/category"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockCategoryRepo struct{}

var (
	Create func(category *category.Category) error
	List   func() ([]category.Category, error)
)

func (m *mockCategoryRepo) Create(category *category.Category) error {
	return Create(category)
}
func (m *mockCategoryRepo) List() ([]category.Category, error) {
	return List()
}

var categoryAppMock category.CategoryService = &mockCategoryRepo{}

func Test_CreateCategory(t *testing.T) {
	Create = func(category *category.Category) error {
		return nil
	}
	var category = &category.Category{
		Type: "test",
	}

	err := categoryAppMock.Create(category)
	assert.Nil(t, err)
}

func Test_ListCategory(t *testing.T) {
	List = func() ([]category.Category, error) {
		return []category.Category{}, nil
	}

	_, err := categoryAppMock.List()
	assert.Nil(t, err)
}
