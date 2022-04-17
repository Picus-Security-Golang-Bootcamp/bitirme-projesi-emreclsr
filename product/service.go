package product

import (
	"fmt"
	"github.com/emreclsr/picusfinal/pagination"
	"go.uber.org/zap"
)

type productService struct {
	repo ProductRepository
}

type ProductService interface {
	Create(product *Product) error
	Search(word string) ([]Product, error)
	Delete(id uint) error
	Update(product *Product) error
	List(pg *pagination.Pagination) (*pagination.Pagination, error)
	Get(id uint) (*Product, error)
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

// Compile time proof of interface implementation
var _ ProductService = &productService{}

func (s *productService) Create(product *Product) error {
	zap.L().Info("Create product service triggered")
	err := s.repo.Create(product)
	if err != nil {
		zap.L().Error("Error creating product (service)", zap.Error(err))
		return err
	}
	return nil
}

func (s *productService) Search(word string) ([]Product, error) {
	zap.L().Info("Search product service triggered")
	products, err := s.repo.Search(word)
	if err != nil {
		zap.L().Error("Error searching product (service)", zap.Error(err))
		return nil, err
	}
	return products, nil
}

func (s *productService) Delete(id uint) error {
	zap.L().Info("Delete product service triggered")
	err := s.repo.Delete(id)
	if err != nil {
		zap.L().Error("Error deleting product (service)", zap.Error(err))
		return err
	}
	return nil
}

func (s *productService) Update(product *Product) error {
	zap.L().Info("Update product service triggered")
	err := s.repo.Update(product)
	if err != nil {
		zap.L().Error("Error updating product (service)", zap.Error(err))
		return err
	}
	return nil
}

func (s *productService) List(pg *pagination.Pagination) (*pagination.Pagination, error) {
	operationResult, err := s.repo.List(pg)

	if err != nil {
		return nil, err
	}

	var data = operationResult

	//set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", pg.URLPath, pg.Limit, 0, pg.Sort)
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", pg.URLPath, pg.Limit, pg.TotalPages, pg.Sort)

	if data.Page > 0 {
		//set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", pg.URLPath, pg.Limit, data.Page-1, pg.Sort)
	}
	if data.Page < pg.TotalPages {
		//set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", pg.URLPath, pg.Limit, data.Page+1, pg.Sort)
	}

	if data.Page > pg.TotalPages {
		//reset previous page pagination response
		data.PreviousPage = ""
	}
	return data, nil

}

func (s *productService) Get(id uint) (*Product, error) {
	zap.L().Info("Get product service triggered")
	product, err := s.repo.Get(id)
	if err != nil {
		zap.L().Error("Error getting product (service)", zap.Error(err))
		return nil, err
	}
	return product, nil
}
