package category_test

import (
	"bytes"
	"fmt"
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/category"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/handlers"
	"github.com/emreclsr/picusfinal/product"
	"github.com/emreclsr/picusfinal/repositories"
	"github.com/emreclsr/picusfinal/services"
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func init() {
	er := godotenv.Load("./../.env")
	if er != nil {
		fmt.Println("Error loading .env file")
	}
}

func TestGetAllCategories(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&category.Category{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	app := gin.Default()
	app.GET("/category", hands.Category.GetAllCategories)
	err = category.NewCategoryRepository(DB).Create(&category.Category{Type: "test", Product: []product.Product{}})
	assert.Nil(t, err)

	req := httptest.NewRequest("GET", "/category", nil)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestCreateCategoryFromCSV(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{}, &category.Category{}, &product.Product{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	token := authentication.NewToken()
	repos := repositories.NewRepositories(DB)
	servs := services.NewServices(DB, *repos)
	hands := handlers.NewHandlers(*servs, token)

	app := gin.Default()
	app.POST("/category", hands.Category.CreateCategoryFromCSV)
	w := httptest.NewRecorder()

	cookie := &http.Cookie{
		Name:    "TokenJWT",
		Value:   db.Hastoken,
		Expires: time.Now().Add(time.Hour * 5),
	}
	file, _ := os.Open("testdummy.csv")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("csv", file.Name())
	io.Copy(part, file)
	writer.Close()
	req, _ := http.NewRequest("POST", "/category", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.AddCookie(cookie)
	app.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
