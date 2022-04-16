package basket

import (
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/user"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{}, &Basket{})
	db.AddUser(DB)
	defer db.DropDB(DB)

	repo := NewBasketRepository(DB)
	err = repo.Create(1)
	assert.Nil(t, err)
}

func Test_Update_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &Basket{})
	assert.Nil(t, err)
	db.AddUser(DB)
	defer db.DropDB(DB)

	repo := NewBasketRepository(DB)
	var basket = &Basket{
		UserID:     1,
		ProductIds: pq.Int64Array{1, 2, 3},
		Amount:     pq.Int64Array{1, 2, 3},
		TotalPrice: 999,
	}
	err = repo.Update(basket)
	assert.Nil(t, err)
}

func Test_GetByUserId(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &Basket{})
	assert.Nil(t, err)
	db.AddUser(DB)
	defer db.DropDB(DB)

	repo := NewBasketRepository(DB)

	err = repo.Create(1)
	assert.Nil(t, err)

	bskt, err := repo.GetByUserId(uint(1))
	assert.Nil(t, err)
	assert.Equal(t, uint(1), bskt.UserID)
}
