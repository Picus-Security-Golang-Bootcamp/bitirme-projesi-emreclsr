package basket_test

import (
	"fmt"
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

func TestGetBasket(t *testing.T) {
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
	app.GET("/basket", hands.Basket.GetBasket)
	assert.Nil(t, err)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}
	req := httptest.NewRequest("GET", "/basket", nil)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestUpdateBasket(t *testing.T) {
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
	app.PUT("/basket", hands.Basket.UpdateBasket)
	assert.Nil(t, err)

	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}
	bodyReader := strings.NewReader(`{"user_id":1, "product_ids":[1], "amount" : [5]}`)
	req := httptest.NewRequest("PUT", "/basket", bodyReader)
	rr := httptest.NewRecorder()
	req.AddCookie(cookie)
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)

}

func TestCreateAnOrder(t *testing.T) {
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

	testBasket := basket.Basket{UserID: 1, ProductIds: []int64{1}, Amount: []int64{5}}
	basketRepo := basket.NewBasketRepository(DB)
	err = basketRepo.Update(&testBasket)
	assert.Nil(t, err)
	app := gin.Default()
	app.POST("/basket", hands.Basket.CreateAnOrder)
	assert.Nil(t, err)

	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}

	req := httptest.NewRequest("POST", "/basket", nil)
	rr := httptest.NewRecorder()
	req.AddCookie(cookie)
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)
}
