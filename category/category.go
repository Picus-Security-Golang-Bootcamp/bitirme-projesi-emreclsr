package category

import (
	"github.com/emreclsr/picusfinal/product"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Type     string            `json:"type"`
	IsDelete bool              `json:"isDelete"`
	Product  []product.Product `json:"product" gorm:"-"`
}
