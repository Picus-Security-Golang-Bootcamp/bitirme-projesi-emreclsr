package authentication

import (
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Authenticate struct {
	userService user.UserService
	token       TokenInterface
}
type IAuthenticate interface {
	Login(c *gin.Context)
}

// NewAuthenticate creates a new Authenticate object (constructor)
func NewAuthenticate(usrSrv user.UserService, tk TokenInterface) IAuthenticate {
	return &Authenticate{userService: usrSrv, token: tk}
}

var _ IAuthenticate = &Authenticate{}
var validate *validator.Validate

// Login godoc
// @Summary Login
// @Description Login
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param body body user.User true "User"
// @Success 200 "Logged in successfully"
// @Failure 401 "Username or password is incorrect"
// @Failure 422 "Invalid JSON provided for login"
// @Failure 500 "Error getting user while logging in"
// @Router /login [post]
func (au *Authenticate) Login(c *gin.Context) {
	zap.L().Info("Login triggered")
	var usr user.User
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		zap.L().Error("Error binding JSON while logging in")
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided for login")
		return
	}
	validate = validator.New()
	err = validate.Struct(&usr)
	if err != nil {
		zap.L().Error("Error in json which sended data to create product handler", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, usrErr := au.userService.GetByEmail(usr.Email)
	if usrErr != nil {
		zap.L().Error("Error getting user while logging in")
		c.JSON(http.StatusInternalServerError, usrErr)
		return
	}
	// Password check
	errPass := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(usr.Password))
	if errPass != nil {
		zap.L().Error("Error comparing password while logging in")
		c.JSON(http.StatusUnauthorized, "Username or password is incorrect")
		return
	}
	// Email check
	if usr.Email != u.Email {
		zap.L().Error("Error comparing email while logging in")
		c.JSON(http.StatusUnauthorized, "Username or password is incorrect")
		return
	}
	tkstring, tknErr := au.token.CreateToken(u)
	if tknErr != nil {
		zap.L().Error("Error creating token while logging in")
		c.JSON(http.StatusInternalServerError, tknErr)
		return
	}
	c.SetCookie("TokenJWT", tkstring, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "Logged in successfully")
}
