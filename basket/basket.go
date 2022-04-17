package basket

import (
	"errors"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/user"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type Basket struct {
	gorm.Model
	User       user.User         `json:"user"`
	UserID     uint              `json:"user_id" example:"1"`
	ProductIds pq.Int64Array     `json:"product_ids" gorm:"type:integer[]" example:"1,2,3"`
	Products   []product.Product `json:"products" gorm:"-"`
	Amount     pq.Int64Array     `json:"amount" gorm:"type:integer[]" example:"1,2,3"`
	TotalPrice float64           `json:"total_price" example:"999.99"`
}

func (b *Basket) CalculateTotalPrice() {
	for i, v := range b.Amount {
		b.TotalPrice += float64(v) * b.Products[i].Price
	}
}

func (b *Basket) CalculateLineTotal() pq.Float64Array {
	var lineTotal []float64
	for i, v := range b.Amount {
		lineTotal = append(lineTotal, float64(v)*b.Products[i].Price)
	}
	return pq.Float64Array(lineTotal)
}

func (b *Basket) ToOrder() *order.Order {
	o := order.Order{
		UserID:     b.UserID,
		User:       b.User,
		TotalPrice: b.TotalPrice,
		IsCanceled: false,
		Amount:     b.Amount,
		LineTotal:  b.CalculateLineTotal(),
		ProductIds: b.ProductIds,
		Products:   b.Products,
	}
	return &o
}

func (b *Basket) CheckItemsCountAndBasketQuantity() (bool, error) {
	maxItem, err := strconv.Atoi(os.Getenv("MAX_ITEM_PER_BASKET"))
	if err != nil {
		zap.L().Error("Error while converting MAX_ITEM_PER_BASKET to int", zap.Error(err))
		return false, err
	}
	if len(b.ProductIds) > maxItem {

		return false, errors.New("max item limit exceeded")
	}

	maxAmount, err := strconv.Atoi(os.Getenv("MAX_QTY_PER_PRODUCT"))
	if err != nil {
		zap.L().Error("Error while converting MAX_AMOUNT_OF_ITEM to int", zap.Error(err))
		return false, err
	}
	for _, v := range b.Amount {
		if v > int64(maxAmount) {
			return false, errors.New("max amount limit exceeded")
		}
	}
	return true, nil
}
