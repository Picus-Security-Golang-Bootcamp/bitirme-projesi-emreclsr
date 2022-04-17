package user

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name" example:"Test User"`
	Email    string `json:"email" example:"test@tst.com" validate:"required,email"`
	Password string `json:"password" example:"password" validate:"required"`
	Phone    string `json:"phone" example:"+05341234567"`
	Address  string `json:"address" example:"Test Address"`
	Role     string `json:"role" example:"admin" `
	Status   string `json:"status" example:"active"`
}

// BeforeCreate will be hashed password before create User
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("Error hashing password", zap.Error(err))
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
