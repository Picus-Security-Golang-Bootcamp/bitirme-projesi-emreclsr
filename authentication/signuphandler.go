package authentication

import (
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Users struct {
	userServ user.UserService
}

// NewUsers creates a new Users object (constructor)
func NewUsers(usrService user.UserService) *Users {
	return &Users{userServ: usrService}
}

func (u *Users) SignUp(c *gin.Context) {
	var usr user.User

	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided for signup")
		return
	}
	check, _ := u.userServ.GetByEmail(usr.Email)
	if check != nil {
		c.JSON(http.StatusConflict, "User already exists")
		return
	}

	err = u.userServ.Create(&usr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	c.JSON(http.StatusCreated, usr)
}
