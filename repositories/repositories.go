package repositories

import (
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/user"
	"gorm.io/gorm"
)

// Repositories struct collect all repositories into one struct
type Repositories struct {
	DB       *gorm.DB
	Basket   basket.BasketRepository
	Category category.CategoryRepository
	Order    order.OrderRepository
	Product  product.ProductRepository
	User     user.UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		DB:       db,
		Basket:   basket.NewBasketRepository(db),
		Category: category.NewCategoryRepository(db),
		Order:    order.NewOrderRepository(db),
		Product:  product.NewProductRepository(db),
		User:     user.NewUserRepository(db),
	}
}
