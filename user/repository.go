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
	zap.L().Info("Creating user")
	err := r.db.Create(&user).Error
	if err != nil {
		zap.L().Error("Error creating user", zap.Error(err))
		return err
	}
	return nil
}

func (r *repository) GetByEmail(email string) (*User, error) {
	zap.L().Info("Getting user by email", zap.String("email", email))
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetByID(id uint) (*User, error) {
	zap.L().Info("Getting user by id", zap.Uint("id", id))
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		zap.L().Error("Error getting user by id", zap.Error(err))
		return nil, err
	}
	return &user, nil
}
