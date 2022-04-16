package basket_test

import (
	"fmt"
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	er := godotenv.Load("./../.env")
	if er != nil {
		fmt.Println("Error loading .env file")
	}
}

func TestCalculateTotalPrice(t *testing.T) {
	tests := []struct {
		name   string
		Basket basket.Basket
		want   float64
	}{
		{name: "test1", Basket: basket.Basket{}, want: 0},
		{name: "test2", Basket: basket.Basket{Amount: pq.Int64Array{1, 2}, Products: []product.Product{{Price: 10}, {Price: 20}}}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var basket = tt.Basket
			basket.CalculateTotalPrice()
			assert.Equal(t, tt.want, basket.TotalPrice)
		})
	}
}

func TestCalculateLineTotal(t *testing.T) {
	tests := []struct {
		name   string
		Basket basket.Basket
		want   pq.Float64Array
	}{
		{name: "test1", Basket: basket.Basket{}, want: pq.Float64Array(nil)},
		{name: "test2", Basket: basket.Basket{Amount: pq.Int64Array{1, 2}, Products: []product.Product{{Price: 10}, {Price: 20}}}, want: pq.Float64Array{10, 40}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var basket = tt.Basket
			assert.Equal(t, tt.want, basket.CalculateLineTotal())
		})
	}
}

func TestToOrder(t *testing.T) {
	tests := []struct {
		name   string
		basket basket.Basket
		want   *order.Order
	}{
		{name: "test1", basket: basket.Basket{}, want: &order.Order{}},
		{name: "test2",
			basket: basket.Basket{Amount: pq.Int64Array{1, 2}, ProductIds: pq.Int64Array{1, 2}, Products: []product.Product{{Price: 10}, {Price: 20}}, TotalPrice: 50},
			want:   &order.Order{TotalPrice: 50, IsCanceled: false, Amount: pq.Int64Array{1, 2}, LineTotal: pq.Float64Array{10, 40}, ProductIds: pq.Int64Array{1, 2}, Products: []product.Product{{Price: 10}, {Price: 20}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bskt = tt.basket
			assert.Equal(t, tt.want, bskt.ToOrder())
		})
	}
}

func TestCheckItemsCountAndBasketQuantity(t *testing.T) {

	tests :=
		[]struct {
			name   string
			basket basket.Basket
			want   bool
		}{
			{name: "test1", basket: basket.Basket{}, want: true},
			{name: "test2", basket: basket.Basket{ProductIds: pq.Int64Array{1, 2}, Amount: pq.Int64Array{1, 2}}, want: true},
			{name: "test3", basket: basket.Basket{ProductIds: pq.Int64Array{1, 2, 3, 4, 5, 6}, Amount: pq.Int64Array{1, 1, 1, 1, 1, 1}}, want: false},
			{name: "test4", basket: basket.Basket{ProductIds: pq.Int64Array{1, 2, 3, 4}, Amount: pq.Int64Array{1, 1, 1, 99}}, want: false},
			{name: "test5", basket: basket.Basket{ProductIds: pq.Int64Array{1, 2, 3, 4, 5, 6}, Amount: pq.Int64Array{1, 1, 1, 1, 1, 99}}, want: false},
		}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var basket = tt.basket
			b, _ := basket.CheckItemsCountAndBasketQuantity()
			assert.Equal(t, tt.want, b)
		})
	}

}
