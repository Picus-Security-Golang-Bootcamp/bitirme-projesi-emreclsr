package user

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

// Compile time proof of interface implementation
var _ UserRepository = &repository{}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	zap.L().Debug("Creating user triggered")
	err := r.db.Create(&user).Error
	if err != nil {
		zap.L().Error("Error creating user", zap.Error(err))
		return err
	}
	return nil
}

func (r *repository) GetByEmail(email string) (*User, error) {
	zap.L().Debug("Getting user by email")
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetByID(id uint) (*User, error) {
	zap.L().Debug("Getting user by id")
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		zap.L().Error("Error getting user by id")
		return nil, err
	}
	return &user, nil
}
