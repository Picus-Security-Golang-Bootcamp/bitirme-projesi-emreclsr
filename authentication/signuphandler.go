package authentication

import (
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

type Users struct {
	userServ user.UserService
}
type IUsers interface {
	SignUp(c *gin.Context)
}

var _ IUsers = &Users{}

// NewUsers creates a new Users object (constructor)
func NewUsers(usrService user.UserService) IUsers {

	return &Users{userServ: usrService}
}

// SignUp godoc
// @Summary SignUp
// @Description SignUp
// @Tags NewUser
// @Accept  json
// @Produce  json
// @Param body body user.User true "User"
// @Success 200 {object} user.User
// @Failure 409 "User already exists"
// @Failure 422 "Invalid JSON provided for signup"
// @Failure 500 "Server error"
// @Router /user [post]
func (u *Users) SignUp(c *gin.Context) {
	zap.L().Info("SignUp triggered")
	var usr user.User

	err := c.ShouldBindJSON(&usr)
	if err != nil {
		zap.L().Error("Error binding JSON while signing up")
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON provided for signup")
		return
	}
	validate = validator.New()
	err = validate.Struct(&usr)
	if err != nil {
		zap.L().Error("Error in json which sended data to create product handler", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	check, _ := u.userServ.GetByEmail(usr.Email)
	if check != nil {
		zap.L().Error("Error while getting email in sing up")
		c.JSON(http.StatusConflict, "User already exists")
		return
	}

	err = u.userServ.Create(&usr)
	if err != nil {
		zap.L().Error("Error signing up create")
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	c.JSON(http.StatusCreated, usr)
}
