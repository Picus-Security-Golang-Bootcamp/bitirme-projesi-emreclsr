package product

import (
	"go.uber.org/zap"
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
	zap.L().Info("Create product", zap.Reflect("product", product))
	err := r.db.Create(&product).Error
	if err != nil {
		zap.L().Error("Create product error (repository)", zap.Error(err))
		return err
	}
	return nil
}

func (r *repository) Search(word string) ([]Product, error) {
	zap.L().Info("Searching for products", zap.String("word", word))
	var products []Product
	err := r.db.Where("name LIKE ?", "%"+word+"%").Find(&products).Error
	if err != nil {
		zap.L().Error("Searching for products error (repository)", zap.Error(err))
		return nil, err
	}
	return products, nil
}

func (r *repository) Delete(id uint) error {
	zap.L().Info("Deleting product", zap.Uint("id", id))
	err := r.db.Where("id = ?", id).Delete(&Product{}).Error
	if err != nil {
		zap.L().Error("Deleting product error (repository)", zap.Error(err))
		return err
	}
	return nil
}

func (r *repository) Update(product *Product) error {
	zap.L().Info("Updating product", zap.Reflect("product", product))
	err := r.db.Save(&product).Error
	if err != nil {
		zap.L().Error("Updating product error (repository)", zap.Error(err))
		return err
	}
	return nil
}

func (r *repository) List() ([]Product, error) {
	zap.L().Info("Listing products", zap.Reflect("products", []Product{}))
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		zap.L().Error("Listing products error (repository)", zap.Error(err))
		return nil, err
	}
	return products, nil
}

func (r *repository) Get(id uint) (*Product, error) {
	zap.L().Info("Getting product", zap.Uint("id", id))
	product := &Product{}
	product.ID = id
	err := r.db.First(&product).Error
	if err != nil {
		zap.L().Error("Getting product error (repository)", zap.Error(err))
		return nil, err
	}
	return product, nil
}
