package user

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
	err := s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetByEmail(email string) (*User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
