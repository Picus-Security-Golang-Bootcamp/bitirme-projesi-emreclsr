package order

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	Create(order *Order) error
	Get(id uint) (*Order, error)
	List(userID uint) ([]Order, error)
	Update(order *Order) error
}

type repository struct {
	db *gorm.DB
}

// Compile time proof of interface implementation
var _ OrderRepository = &repository{}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) Create(order *Order) error {
	err := r.db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(id uint) (*Order, error) {
	var order Order
	err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *repository) List(userID uint) ([]Order, error) {
	var orders []Order
	err := r.db.Preload(clause.Associations).Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *repository) Update(order *Order) error {
	err := r.db.Save(order).Error
	if err != nil {
		return err
	}
	return nil
}
