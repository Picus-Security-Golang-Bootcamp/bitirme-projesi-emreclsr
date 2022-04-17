package order

import (
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/user"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     uint              `json:"user_id" example:"1"`
	User       user.User         `json:"user" gorm:"foreignkey:UserID"`
	TotalPrice float64           `json:"total_price" example:"999.99"`
	IsCanceled bool              `json:"is_canceled" example:"false"`
	Amount     pq.Int64Array     `json:"amount" gorm:"type:integer[]" example:"1,2,3"`
	LineTotal  pq.Float64Array   `json:"line_total" gorm:"type:float[]" example:"1.99,2.99,3.99"`
	ProductIds pq.Int64Array     `json:"product_ids" gorm:"type:integer[]" example:"1,2,3"`
	Products   []product.Product `json:"products" gorm:"many2many:order_products;association_foreignkey:ID;foreignkey:ID"`
}

func (o *Order) CheckTime() bool {
	duration := time.Duration(14 * 24 * time.Hour)
	if o.CreatedAt.Add(duration).After(time.Now()) {
		return true
	}
	return false
}
