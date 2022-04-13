package user

import "go.uber.org/zap"

type userService struct {
	repo UserRepository
}

type UserService interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id uint) (*User, error)
}

// Compile time proof of interface implementation
var _ UserService = &userService{}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *User) error {
	zap.L().Info("Create user triggered", zap.String("email", user.Email))
	err := s.repo.Create(user)
	if err != nil {
		zap.L().Error("Error creating user", zap.Error(err))
		return err
	}
	return nil
}

func (s *userService) GetByEmail(email string) (*User, error) {
	zap.L().Info("Get user by email triggered", zap.String("email", email))
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		//zap.L().Error("Error getting user by email")
		return nil, err
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*User, error) {
	zap.L().Info("Get user by id triggered", zap.Uint("id", id))
	user, err := s.repo.GetByID(id)
	if err != nil {
		zap.L().Error("Error getting user by id", zap.Error(err))
		return nil, err
	}
	return user, nil
}
