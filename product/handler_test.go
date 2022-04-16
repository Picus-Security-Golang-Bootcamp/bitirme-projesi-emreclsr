package product_test

import (
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/basket"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/handlers"
	"github.com/emreclsr/picusfinal/order"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/repositories"
	"github.com/emreclsr/picusfinal/services"
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateProduct(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	app := gin.Default()
	app.POST("/product", hands.Product.CreateProduct)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}
	bodyReader := strings.NewReader(`{"Name": "test", "Price": 100, "Stock": 10, "Type": "testcategory"}`)
	req := httptest.NewRequest("POST", "/product", bodyReader)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code)
}

func TestGetProduct(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	assert.Nil(t, err)
	defer db.DropDB(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	testProduct := product.Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := product.NewProductRepository(DB)
	err = repo.Create(&testProduct)
	assert.Nil(t, err)

	app := gin.Default()
	app.GET("/product", hands.Product.GetProducts)

	req := httptest.NewRequest("GET", "/product", nil)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestSearch(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	testProduct := product.Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := product.NewProductRepository(DB)
	err = repo.Create(&testProduct)
	assert.Nil(t, err)

	app := gin.Default()
	app.GET("/product/:word", hands.Product.GetProducts)

	req := httptest.NewRequest("GET", "/product/test", nil)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}

func TestDelete(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	testProduct := product.Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := product.NewProductRepository(DB)
	err = repo.Create(&testProduct)

	app := gin.Default()
	app.DELETE("/product/:id", hands.Product.DeleteProduct)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}

	req := httptest.NewRequest("DELETE", "/product/1", nil)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}

func TestUpdateProduct(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{}, &basket.Basket{}, &order.Order{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	testProduct := product.Product{Name: "test", Price: 100, Stock: 10, Type: "testcategory"}
	repo := product.NewProductRepository(DB)
	err = repo.Create(&testProduct)

	app := gin.Default()
	app.PUT("/product/:id", hands.Product.UpdateProduct)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}

	bodyReader := strings.NewReader(`{"Name": "test", "Price": 500, "Stock": 10, "Type": "testcategory"}`)
	req := httptest.NewRequest("PUT", "/product/1", bodyReader)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}
