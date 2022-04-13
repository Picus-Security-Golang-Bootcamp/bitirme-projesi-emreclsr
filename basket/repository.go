package basket

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

type BasketRepository interface {
	Create(userId uint) error
	Update(basket *Basket) error
	GetByUserId(userId uint) (*Basket, error)
}

// Compile time proof of interface implementation
var _ BasketRepository = &repository{}

func NewBasketRepository(db *gorm.DB) BasketRepository {
	return &repository{db: db}
}

func (r *repository) Create(userId uint) error {
	zap.L().Info("Create basket (repository)")
	basket := Basket{UserID: userId}
	return r.db.Preload(clause.Associations).Create(&basket).Error
}

func (r *repository) Update(basket *Basket) error {
	zap.L().Info("Update basket (repository)")
	return r.db.Preload(clause.Associations).Where("user_id = ?", basket.UserID).Save(&basket).Error
}

func (r *repository) GetByUserId(userId uint) (*Basket, error) {
	zap.L().Info("Get basket by user id (repository)")
	basket := &Basket{}
	err := r.db.Preload(clause.Associations).First(&basket).Error
	return basket, err
}
