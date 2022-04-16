package product_test

import (
	"github.com/emreclsr/picusfinal/pagination"
	"github.com/emreclsr/picusfinal/product"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockProductRepo struct{}

var (
	Create func(prd *product.Product) error
	Search func(word string) ([]product.Product, error)
	Delete func(id uint) error
	Update func(prd *product.Product) error
	List   func(pg *pagination.Pagination) (*pagination.Pagination, error)
	Get    func(id uint) (*product.Product, error)
)

func (m *mockProductRepo) Create(product *product.Product) error {
	return Create(product)
}
func (m *mockProductRepo) Search(word string) ([]product.Product, error) {
	return Search(word)
}
func (m *mockProductRepo) Delete(id uint) error {
	return Delete(id)
}
func (m *mockProductRepo) Update(product *product.Product) error {
	return Update(product)
}
func (m *mockProductRepo) List(pg *pagination.Pagination) (*pagination.Pagination, error) {
	return List(pg)
}
func (m *mockProductRepo) Get(id uint) (*product.Product, error) {
	return Get(id)
}

var productAppMock product.ProductService = &mockProductRepo{}

func Test_Create(t *testing.T) {
	Create = func(prd *product.Product) error {
		return nil
	}
	var prd product.Product
	prd.Name = "test"
	prd.Price = 10
	prd.Stock = 10
	prd.Type = "testCategory"

	err := productAppMock.Create(&prd)
	assert.Nil(t, err)
}

func Test_Search(t *testing.T) {
	Search = func(word string) ([]product.Product, error) {
		return []product.Product{}, nil
	}
	_, err := productAppMock.Search("test")
	assert.Nil(t, err)
}

func Test_Delete(t *testing.T) {
	Delete = func(id uint) error {
		return nil
	}
	err := productAppMock.Delete(1)
	assert.Nil(t, err)
}

func Test_Update(t *testing.T) {
	Update = func(prd *product.Product) error {
		return nil
	}
	var prd product.Product
	prd.Name = "test"
	prd.Price = 10
	prd.Stock = 10
	prd.Type = "testCategory"

	err := productAppMock.Update(&prd)
	assert.Nil(t, err)
}

func Test_List(t *testing.T) {
	List = func(pg *pagination.Pagination) (*pagination.Pagination, error) {
		return &pagination.Pagination{}, nil
	}
	_, err := productAppMock.List(&pagination.Pagination{})
	assert.Nil(t, err)
}

func Test_Get(t *testing.T) {
	Get = func(id uint) (*product.Product, error) {
		return &product.Product{}, nil
	}
	_, err := productAppMock.Get(1)
	assert.Nil(t, err)
}
