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
	Basket         basket.BasketHandler
	Category       category.CategoryHandler
	Order          order.OrderHandler
	Product        product.ProductHandler
	User           authentication.Users
	Authentication authentication.Authenticate
}

func NewHandlers(services services.Services, token authentication.Token) *Handlers {
	return &Handlers{
		Basket:         *basket.NewBasketHandler(services.Basket, token, services.User, services.Order, services.Product),
		Category:       *category.NewCategoryHandler(services.Category, token, services.Product),
		Order:          *order.NewOrderHandler(services.Order, token),
		Product:        *product.NewProductHandler(services.Product, token),
		User:           *authentication.NewUsers(services.User),
		Authentication: *authentication.NewAuthenticate(services.User, &token),
	}
}
