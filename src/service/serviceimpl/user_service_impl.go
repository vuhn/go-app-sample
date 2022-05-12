package serviceimpl

import (
	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/errs"
	"github.com/vuhn/go-app-sample/infrastructure/repository"
	"github.com/vuhn/go-app-sample/service"
	"golang.org/x/crypto/bcrypt"
)

// NewUserService return implementation of user service interface
func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

// UserServiceImpl implements UserService interface
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// CreateUser creates take user entity and create an user
func (u *UserServiceImpl) CreateUser(user *entity.User) (*entity.User, error) {
	existedUser, err := u.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	if existedUser != nil {
		return nil, errs.ErrEmailExisted
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	user.Password = string(encryptedPassword)
	userEntity, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	return userEntity, nil
}

// Login is method to very user credentials and return a access token
func (u *UserServiceImpl) Login(email string, password string) (string, error) {
	// TODO: Implement login method
	panic("not implemented")
}
