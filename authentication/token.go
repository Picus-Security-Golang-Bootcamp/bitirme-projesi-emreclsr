package authentication

import (
	"fmt"
	"github.com/emreclsr/picusfinal/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type Token struct {
	UserID uint
	Role   string
	jwt.StandardClaims
}

type TokenInterface interface {
	CreateToken(usr *user.User) (string, error)
	VerifyToken(r *gin.Context) (*Token, error)
}

// Compile time proof of interface implementation
var _ TokenInterface = &Token{}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) CreateToken(usr *user.User) (string, error) {
	token := &Token{
		UserID: usr.ID,
		Role:   usr.Role,
	}
	token.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	at, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return at, nil
}

func (t *Token) VerifyToken(r *gin.Context) (*Token, error) {
	var claim = &Token{}
	cookie, err := r.Request.Cookie("TokenJWT")
	if err != nil {
		return nil, err
	}
	tokenStr := cookie.Value
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}
	return claim, nil
}
