package handlers

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/services"
)

type Handlers struct {
	Basket         basket.IBasketHandler
	Category       category.ICategoryHandler
	Order          order.IOrderHandler
	Product        product.IProductHandler
	User           authentication.IUsers
	Authentication authentication.IAuthenticate
}

func NewHandlers(services services.Services, token authentication.TokenInterface) *Handlers {
	return &Handlers{
		Basket:         basket.NewBasketHandler(services.Basket, token, services.User, services.Order, services.Product),
		Category:       category.NewCategoryHandler(services.Category, token, services.Product),
		Order:          order.NewOrderHandler(services.Order, token),
		Product:        product.NewProductHandler(services.Product, token),
		User:           authentication.NewUsers(services.User),
		Authentication: authentication.NewAuthenticate(services.User, token),
	}
}
