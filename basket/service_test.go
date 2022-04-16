package basket_test

import (
	"github.com/emreclsr/picusfinal/basket"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockBasketRepo struct{}

var (
	Create      func(userId uint) error
	Update      func(basket *basket.Basket) error
	GetByUserId func(userId uint) (*basket.Basket, error)
)

func (m *mockBasketRepo) CreateBasket(userId uint) error {
	return Create(userId)
}
func (m *mockBasketRepo) UpdateBasket(basket *basket.Basket) error {
	return Update(basket)
}
func (m *mockBasketRepo) GetByUserId(userId uint) (*basket.Basket, error) {
	return GetByUserId(userId)
}

var basketAppMock basket.BasketService = &mockBasketRepo{}

func Test_CreateBasket(t *testing.T) {
	Create = func(userId uint) error {
		return nil
	}
	err := basketAppMock.CreateBasket(1)
	assert.Nil(t, err)
}

func Test_UpdateBasket(t *testing.T) {
	Update = func(basket *basket.Basket) error {
		return nil
	}
	var basket = &basket.Basket{
		UserID:     1,
		ProductIds: pq.Int64Array{1, 2, 3},
		Amount:     pq.Int64Array{1, 2, 3},
		TotalPrice: 999,
	}
	err := basketAppMock.UpdateBasket(basket)
	assert.Nil(t, err)
}

func Test_GetBasketByUserId(t *testing.T) {
	GetByUserId = func(userId uint) (*basket.Basket, error) {
		return &basket.Basket{
			UserID:     1,
			ProductIds: pq.Int64Array{1, 2, 3},
			Amount:     pq.Int64Array{1, 2, 3},
			TotalPrice: 999,
		}, nil
	}
	basket, err := basketAppMock.GetByUserId(1)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), basket.UserID)
}
