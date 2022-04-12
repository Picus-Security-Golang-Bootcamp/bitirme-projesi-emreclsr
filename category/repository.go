package category

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	catType := category.Type
	rowsAffected := r.db.Where("type = ?", catType).Updates(&category).RowsAffected
	if rowsAffected == 0 {
		err := r.db.Create(category).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) List() ([]Category, error) {
	var categories []Category
	//Find: get all IsDelete false rows
	err := r.db.Preload(clause.Associations).Find(&categories, "is_delete = false").Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
