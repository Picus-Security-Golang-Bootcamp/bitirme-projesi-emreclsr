package services

import (
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/repositories"
	"github.com/emreclsr/picusfinal/user"
	"gorm.io/gorm"
)

// Services struct collect all services into one struct
type Services struct {
	Basket   basket.BasketService
	Category category.CategoryService
	Order    order.OrderService
	Product  product.ProductService
	User     user.UserService
}

func NewServices(db *gorm.DB, repo repositories.Repositories) *Services {
	return &Services{
		Basket:   basket.NewBasketService(repo.Basket),
		Category: category.NewCategoryService(repositories.NewRepositories(db).Category),
		Order:    order.NewOrderService(repositories.NewRepositories(db).Order),
		Product:  product.NewProductService(repositories.NewRepositories(db).Product),
		User:     user.NewUserService(repositories.NewRepositories(db).User),
	}
}
