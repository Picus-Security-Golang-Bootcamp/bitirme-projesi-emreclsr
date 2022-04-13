package category

import (
	"github.com/emreclsr/picusfinal/product"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Type    string            `json:"type"`
	Product []product.Product `json:"product" gorm:"many2many:category_product;association_foreignkey:ID;foreignkey:ID"`
}
