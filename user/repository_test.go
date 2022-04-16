package user_test

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
	DB.AutoMigrate(&user.User{})
	defer db.DropDB(DB)

	testUser := user.User{
		FullName: "test",
		Email:    "test@user.com",
		Password: "test",
		Phone:    "05123456789",
		Address:  "test address",
		Role:     "admin",
		Status:   "active"}

	repo := user.NewUserRepository(DB)
	err = repo.Create(&testUser)
	assert.Nil(t, err)
}

func Test_GetByEmail_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{})
	defer db.DropDB(DB)

	testUser := user.User{
		FullName: "test",
		Email:    "test@user.com",
		Password: "test",
		Phone:    "05123456789",
		Address:  "test address",
		Role:     "admin",
		Status:   "active"}

	repo := user.NewUserRepository(DB)
	err = repo.Create(&testUser)
	assert.Nil(t, err)
	usr, err := repo.GetByEmail("test@user.com")
	assert.Nil(t, err)
	assert.Equal(t, usr.Address, "test address")
	assert.Equal(t, usr.Role, "admin")
}

func Test_GetByID_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&user.User{})
	defer db.DropDB(DB)

	testUser := user.User{
		FullName: "test",
		Email:    "test@user.com",
		Password: "test",
		Phone:    "05123456789",
		Address:  "test address",
		Role:     "admin",
		Status:   "active"}

	repo := user.NewUserRepository(DB)
	err = repo.Create(&testUser)
	assert.Nil(t, err)
	usr, err := repo.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, usr.FullName, "test")
	assert.Equal(t, usr.Email, "test@user.com")
}
