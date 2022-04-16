package product

import (
	"fmt"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/pagination"
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
	DB.AutoMigrate(&Product{})
	defer db.DropDB(DB)

	product := Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := NewProductRepository(DB)
	err = repo.Create(&product)
	assert.Nil(t, err)
}

func Test_Search_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Product{})
	defer db.DropDB(DB)

	product := Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := NewProductRepository(DB)
	err = repo.Create(&product)
	assert.Nil(t, err)
	productList, err := repo.Search("test")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(productList))
}

func Test_Delete_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Product{})
	defer db.DropDB(DB)

	product := Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := NewProductRepository(DB)
	err = repo.Create(&product)
	assert.Nil(t, err)
	err = repo.Delete(product.ID)
	assert.Nil(t, err)

}

func Test_Update_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Product{})
	defer db.DropDB(DB)

	product := Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := NewProductRepository(DB)
	err = repo.Create(&product)
	assert.Nil(t, err)
	product.Name = "test2"
	err = repo.Update(&product)
	assert.Nil(t, err)
	prd, err := repo.Get(product.ID)
	assert.Nil(t, err)
	assert.Equal(t, "test2", prd.Name)
}

func Test_List_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Product{})
	defer db.DropDB(DB)
	var pag pagination.Pagination
	pag.Limit = 10
	product := Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := NewProductRepository(DB)
	err = repo.Create(&product)
	assert.Nil(t, err)
	page, err := repo.List(&pag)

	assert.Nil(t, err)
	assert.Equal(t, 0, page.Page)
	assert.Equal(t, 1, page.TotalRows)
}

func Test_Get_Should_Success(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	DB.AutoMigrate(&Product{})
	defer db.DropDB(DB)

	product := Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := NewProductRepository(DB)
	err = repo.Create(&product)

	p1, err := repo.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), p1.ID)
	assert.Equal(t, "test", p1.Name)
	assert.Equal(t, "testcategory", p1.Type)
}
