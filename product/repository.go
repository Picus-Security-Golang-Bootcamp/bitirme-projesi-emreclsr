package product

import (
	"github.com/emreclsr/picusfinal/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math"
)

type ProductRepository interface {
	Create(product *Product) error
	Search(word string) ([]Product, error)
	Delete(id uint) error
	Update(product *Product) error
	//List() ([]Product, error)
	List(pg *pagination.Pagination) (*pagination.Pagination, error, int)
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
	zap.L().Info("Create product")
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

//func (r *repository) List() ([]Product, error) {
//	zap.L().Info("Listing products")
//	var products []Product
//	err := r.db.Find(&products).Error
//	if err != nil {
//		zap.L().Error("Listing products error (repository)", zap.Error(err))
//		return nil, err
//	}
//	return products, nil
//}

//type RepositoryResult struct {
//	Result interface{}
//	Error  error
//}

func (r *repository) List(pg *pagination.Pagination) (*pagination.Pagination, error, int) {
	zap.L().Info("Listing products")
	var products []Product

	var totalRows int64 = 0
	var totalPages int = 0
	var fromRow int64 = 0
	var toRow int64 = 0

	offset := pg.Page * pg.Limit

	// get data with limit, offset & order
	err := r.db.Limit(pg.Limit).Offset(offset).Order(pg.Sort).Find(&products).Error
	if err != nil {
		zap.L().Error("Listing products error (repository)", zap.Error(err))
		return &pagination.Pagination{}, err, 0
	}
	pg.Rows = products

	// count all data
	err = r.db.Model(&Product{}).Count(&totalRows).Error
	if err != nil {
		zap.L().Error("Counting products error (repository)", zap.Error(err))
		return &pagination.Pagination{}, err, 0
	}
	pg.TotalRows = int(totalRows)

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pg.Limit))) - 1

	if pg.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = int64(pg.Limit)
	} else {
		if pg.Page <= totalPages {
			// calculate from & to row on other pages
			fromRow = int64(pg.Page*pg.Limit + 1)
			toRow = int64((pg.Page + 1) * pg.Limit)
		}
	}
	if toRow > totalRows {
		//set to row with total rows
		toRow = totalRows
	}
	pg.FromRow = int(fromRow)
	pg.ToRow = int(toRow)

	return pg, nil, totalPages
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
