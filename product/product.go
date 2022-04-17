package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" example:"product name" validate:"required"`
	Price float64 `json:"price" example:"99.99" validate:"required,numeric"`
	Stock int     `json:"stock" example:"10" validate:"required,numeric"`
	Type  string  `json:"type" example:"category" validate:"required"`
}
