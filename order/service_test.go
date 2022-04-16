package order_test

import (
	"github.com/emreclsr/picusfinal/order"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockOrderRepo struct{}

var (
	Create func(order *order.Order) error
	Get    func(id uint) (*order.Order, error)
	List   func(userID uint) ([]order.Order, error)
	Update func(order *order.Order) error
)

func (m *mockOrderRepo) Create(order *order.Order) error {
	return Create(order)
}
func (m *mockOrderRepo) Get(id uint) (*order.Order, error) {
	return Get(id)
}
func (m *mockOrderRepo) List(userID uint) ([]order.Order, error) {
	return List(userID)
}
func (m *mockOrderRepo) Update(order *order.Order) error {
	return Update(order)
}

var orderAppMock order.OrderService = &mockOrderRepo{}

func Test_OrderCreate(t *testing.T) {
	Create = func(order *order.Order) error {
		return nil
	}
	var order = &order.Order{
		UserID:     1,
		TotalPrice: 50,
		IsCanceled: false,
		Amount:     pq.Int64Array{1, 2, 3},
		LineTotal:  pq.Float64Array{10, 20, 20},
		ProductIds: pq.Int64Array{1, 2, 3},
	}
	err := orderAppMock.Create(order)
	assert.Nil(t, err)
}

func Test_OrderGet(t *testing.T) {
	type Model struct {
		ID uint `gorm:"primary_key"`
	}

	Get = func(id uint) (*order.Order, error) {
		var order order.Order
		order.ID = 1
		order.UserID = 1
		order.TotalPrice = 50
		order.IsCanceled = false
		return &order, nil
	}
	order, err := orderAppMock.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), order.ID)
}

func Test_OrderList(t *testing.T) {
	List = func(userID uint) ([]order.Order, error) {
		return []order.Order{
			{
				UserID:     1,
				TotalPrice: 50,
				IsCanceled: false,
				Amount:     pq.Int64Array{1, 2, 3},
				LineTotal:  pq.Float64Array{10, 20, 20},
				ProductIds: pq.Int64Array{1, 2, 3},
			},
		}, nil
	}
	orders, err := orderAppMock.List(1)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), orders[0].UserID)
}

func Test_OrderUpdate(t *testing.T) {
	Update = func(order *order.Order) error {
		return nil
	}
	var order = &order.Order{
		UserID:     1,
		TotalPrice: 50,
		IsCanceled: false,
		Amount:     pq.Int64Array{1, 2, 3},
		LineTotal:  pq.Float64Array{10, 20, 20},
		ProductIds: pq.Int64Array{1, 2, 3},
	}
	err := orderAppMock.Update(order)
	assert.Nil(t, err)
}
