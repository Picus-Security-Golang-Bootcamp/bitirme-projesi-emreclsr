package category

import (
	"github.com/emreclsr/picusfinal/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Category{})
	db.AddUser(DB)
	defer db.DropDB(DB)

	repo := NewCategoryRepository(DB)
	err = repo.Create(&Category{Type: "test"})
	assert.Nil(t, err)
}

func Test_List_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Category{})
	defer db.DropDB(DB)

	repo := NewCategoryRepository(DB)
	err = repo.Create(&Category{Type: "test"})
	assert.Nil(t, err)
	err = repo.Create(&Category{Type: "test2"})
	assert.Nil(t, err)
	catList, err := repo.List()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(catList))
}
