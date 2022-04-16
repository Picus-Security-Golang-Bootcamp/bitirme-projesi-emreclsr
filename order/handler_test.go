package order_test

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
	"testing"
	"time"
)

func TestGetOrder(t *testing.T) {
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
	app.GET("/order", hands.Order.GetOrders)
	assert.Nil(t, err)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}

	req := httptest.NewRequest("GET", "/order", nil)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestCancelOrder(t *testing.T) {

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
	var testorder order.Order
	testorder.UserID = 1
	repos.Order.Create(&testorder)
	app := gin.Default()
	app.PUT("/order/:id", hands.Order.CancelOrder)
	assert.Nil(t, err)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}

	req := httptest.NewRequest("PUT", "/order/1", nil)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)
}
func TestOrderCancelShouldReturnFalse(t *testing.T) {
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
	var testorder order.Order
	testorder.UserID = 1
	testorder.CreatedAt = time.Now().Add(-(time.Hour * 500))
	repos.Order.Create(&testorder)
	app := gin.Default()
	app.PUT("/order/:id", hands.Order.CancelOrder)
	assert.Nil(t, err)
	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}

	req := httptest.NewRequest("PUT", "/order/1", nil)
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 400, rr.Code)
}
