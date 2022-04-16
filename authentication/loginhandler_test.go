package authentication_test

import (
	"fmt"
	"github.com/emreclsr/picusfinal/authentication"
	"github.com/emreclsr/picusfinal/db"
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	er := godotenv.Load("./../.env")
	if er != nil {
		fmt.Println("Error loading .env file")
	}
}

func TestLogin(t *testing.T) {
	DB, err := db.DBTestConnect()
	assert.Nil(t, err)
	err = DB.AutoMigrate(&user.User{})
	assert.Nil(t, err)
	defer db.DropDB(DB)
	db.AddUser(DB)

	app := gin.Default()
	app.POST("/login", authentication.NewAuthenticate(user.NewUserService(user.NewUserRepository(DB)), authentication.NewToken()).Login)

	bodyReader := strings.NewReader(`{"email": "test@test.com", "password": "test"}`)
	req := httptest.NewRequest("POST", "/login", bodyReader)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	fmt.Println(rr, req)
	assert.Equal(t, 200, rr.Code)
}
