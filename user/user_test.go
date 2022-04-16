package user_test

import (
	"fmt"
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
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

func Test_BeforeCreateShouldNotEqual(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	assert.Nil(t, err)
	defer db.DropDB(DB)

	testUser := user.User{
		FullName: "test",
		Email:    "test@test.com",
		Password: "test"}

	repo := user.NewUserRepository(DB)
	err = repo.Create(&testUser)
	assert.Nil(t, err)
	u, err := repo.GetByEmail("test@test.com")
	assert.Nil(t, err)
	assert.NotEqual(t, "test", u.Password)

}
