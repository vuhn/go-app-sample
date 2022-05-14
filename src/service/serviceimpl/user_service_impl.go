package serviceimpl

import (
	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/errs"
	"github.com/vuhn/go-app-sample/pkg/password"
	"github.com/vuhn/go-app-sample/pkg/token"
	"github.com/vuhn/go-app-sample/repository"
	"github.com/vuhn/go-app-sample/service"
)

// NewUserService return implementation of user service interface
func NewUserService(
	userRepository repository.UserRepository,
	jwt token.Token,
	password password.Password,
) service.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		token:          jwt,
		password:       password,
	}
}

// UserServiceImpl implements UserService interface
type UserServiceImpl struct {
	userRepository repository.UserRepository
	token          token.Token
	password       password.Password
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

	encryptedPassword, err := u.password.GenerateFromPassword(user.Password)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	user.Password = encryptedPassword
	userEntity, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	return userEntity, nil
}

// Login is method to very user credentials and return a access token
func (u *UserServiceImpl) Login(email string, password string) (string, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", errs.ErrInternalServer
	}
	if user == nil {
		return "", errs.ErrEmailNotFound
	}

	if isValid := u.password.CompareHashAndPassword(user.Password, password); !isValid {
		return "", errs.ErrPasswordInvalid
	}

	// Generate access token
	claims := map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	}
	token, err := u.token.Generate(claims)
	if err != nil {
		return "", errs.ErrInternalServer
	}

	return token, nil
}
