package authentication

import (
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Authenticate struct {
	userService user.UserService
	token       TokenInterface
}

// NewAuthenticate creates a new Authenticate object (constructor)
func NewAuthenticate(usrSrv user.UserService, tk TokenInterface) *Authenticate {
	return &Authenticate{userService: usrSrv, token: tk}
}

func (au *Authenticate) Login(c *gin.Context) {
	var usr user.User
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided for login")
	}
	// TODO: Todo add validator
	u, usrErr := au.userService.GetByEmail(usr.Email)
	if usrErr != nil {
		c.JSON(http.StatusInternalServerError, usrErr)
	}
	// Password check
	errPass := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(usr.Password))
	if errPass != nil {
		c.JSON(http.StatusUnauthorized, "Username or password is incorrect")
	}
	// Email check
	if usr.Email != u.Email {
		c.JSON(http.StatusUnauthorized, "Username or password is incorrect")
	}
	tkstring, tknErr := au.token.CreateToken(u)
	if tknErr != nil {
		c.JSON(http.StatusInternalServerError, tknErr)
	}
	c.SetCookie("TokenJWT", tkstring, 60*60*24, "/", "localhost", false, true)
}
