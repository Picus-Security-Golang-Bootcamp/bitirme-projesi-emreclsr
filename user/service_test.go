package user_test

import (
	"github.com/emreclsr/picusfinal/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockUserRepo struct{}

var (
	Create     func(user *user.User) error
	GetByEmail func(email string) (*user.User, error)
	GetByID    func(id uint) (*user.User, error)
)

func (m *mockUserRepo) Create(u *user.User) error {
	return Create(u)
}
func (m *mockUserRepo) GetByEmail(email string) (*user.User, error) {
	return GetByEmail(email)
}
func (m *mockUserRepo) GetByID(id uint) (*user.User, error) {
	return GetByID(id)
}

var userAppMock user.UserService = &mockUserRepo{}

func Test_Create(t *testing.T) {
	Create = func(u *user.User) error {
		return nil
	}
	var usr user.User
	usr.FullName = "test"
	usr.Email = "test@test.com"
	usr.Password = "test"
	usr.Role = "admin"
	err := userAppMock.Create(&usr)
	assert.Nil(t, err)
}

func Test_GetByEmail(t *testing.T) {
	GetByEmail = func(email string) (*user.User, error) {
		var usr user.User
		usr.ID = 1
		usr.FullName = "test"
		usr.Email = "test@test.com"
		return &usr, nil
	}
	usr, err := userAppMock.GetByEmail("test@test.com")
	assert.Nil(t, err)
	assert.Equal(t, usr.Email, "test@test.com")
}

func Test_GetByID(t *testing.T) {
	GetByID = func(id uint) (*user.User, error) {
		var usr user.User
		usr.ID = 1
		usr.FullName = "test"
		usr.Email = "test@test.com"
		return &usr, nil
	}
	usr, err := userAppMock.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, usr.Email, "test@test.com")
}
