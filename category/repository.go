package category

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *Category) error
	List() ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

// Compile time proof of interface implementation
var _ CategoryRepository = &repository{}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &repository{db: db}
}

func (r *repository) Create(category *Category) error {
	zap.L().Info("Create category (repository)")
	catType := category.Type
	rowsAffected := r.db.Where("type = ?", catType).Updates(&category).RowsAffected
	if rowsAffected == 0 {
		err := r.db.Create(category).Error
		if err != nil {
			zap.L().Error("Create category error (repository)", zap.Error(err))
			return err
		}
	}
	return nil
}

func (r *repository) List() ([]Category, error) {
	zap.L().Info("List categories (repository)")
	var categories []Category
	//Find: get all IsDelete false rows
	err := r.db.Find(&categories, "deleted_at is null").Error
	if err != nil {
		zap.L().Error("List categories error (repository)", zap.Error(err))
		return nil, err
	}
	return categories, nil
}
