package product

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *Product) error
	Search(word string) ([]Product, error)
	Delete(id uint) error
	Update(product *Product) error
	List() ([]Product, error)
	Get(id uint) (*Product, error)
}

type repository struct {
	db *gorm.DB
}

// Compile time proof of interface implementation
var _ ProductRepository = &repository{}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &repository{db: db}
}

func (r *repository) Create(product *Product) error {
	err := r.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Search(word string) ([]Product, error) {
	var products []Product
	err := r.db.Where("name LIKE ?", "%"+word+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) Delete(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(product *Product) error {
	err := r.db.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) List() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) Get(id uint) (*Product, error) {
	product := &Product{}
	product.ID = id
	err := r.db.First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
