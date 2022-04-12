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
	UserID     uint              `json:"user_id"`
	User       user.User         `json:"user" gorm:"foreignkey:UserID"`
	TotalPrice float64           `json:"total_price"`
	IsCanceled bool              `json:"is_canceled"`
	Amount     pq.Int64Array     `json:"amount" gorm:"type:integer[]"`
	LineTotal  pq.Float64Array   `json:"line_total" gorm:"type:float[]"`
	ProductIds pq.Int64Array     `json:"product_ids" gorm:"type:integer[]"`
	Products   []product.Product `json:"products" gorm:"many2many:order_products;association_foreignkey:ID;foreignkey:ID"`
}

func (o *Order) CheckTime() bool {
	if o.CreatedAt.Add(14 * 24 * time.Hour).After(time.Now()) {
		return true
	}
	return false
}
