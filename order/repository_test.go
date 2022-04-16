package order

import (
	"fmt"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/user"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	er := godotenv.Load("./../.env")
	if er != nil {
		fmt.Println("Error loading .env file")
	}
}

func Test_Create_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{}, &Order{})
	db.AddUser(DB)
	defer db.DropDB(DB)
	test_order := Order{UserID: 1, TotalPrice: 50, IsCanceled: false}
	repo := NewOrderRepository(DB)
	err = repo.Create(&test_order)
	assert.Nil(t, err)
}

func Test_Update_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{}, &Order{})
	db.AddUser(DB)
	defer db.DropDB(DB)
	test_order := Order{UserID: 1, TotalPrice: 50, IsCanceled: false}
	repo := NewOrderRepository(DB)
	errCreate := repo.Create(&test_order)
	assert.Nil(t, errCreate)
	err = repo.Update(&test_order)
	assert.Nil(t, err)
}

func Test_Get_By_OrderId_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{}, &Order{})
	db.AddUser(DB)
	defer db.DropDB(DB)
	test_order := Order{UserID: 1, TotalPrice: 50, IsCanceled: false}
	repo := NewOrderRepository(DB)
	err = repo.Create(&test_order)
	assert.Nil(t, err)

	_, err = repo.Get(1)
	assert.Nil(t, err)
}

func Test_List_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{}, &Order{})
	db.AddUser(DB)
	defer db.DropDB(DB)
	db.AddUser(DB)
	test_order1 := Order{UserID: 1, TotalPrice: 50, IsCanceled: false}
	test_order2 := Order{UserID: 1, TotalPrice: 50, IsCanceled: false}
	repo := NewOrderRepository(DB)
	err = repo.Create(&test_order1)
	err2 := repo.Create(&test_order2)
	assert.Nil(t, err)
	assert.Nil(t, err2)

	orders, errList := repo.List(1)
	assert.Nil(t, errList)
	assert.Equal(t, 2, len(orders))
}
