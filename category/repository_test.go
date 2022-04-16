package category_test

import (
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&category.Category{})
	db.AddUser(DB)
	defer db.DropDB(DB)

	repo := category.NewCategoryRepository(DB)
	err = repo.Create(&category.Category{Type: "test"})
	assert.Nil(t, err)
}

func Test_List_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&category.Category{})
	defer db.DropDB(DB)

	repo := category.NewCategoryRepository(DB)
	err = repo.Create(&category.Category{Type: "test"})
	assert.Nil(t, err)
	err = repo.Create(&category.Category{Type: "test2"})
	assert.Nil(t, err)
	catList, err := repo.List()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(catList))
}
