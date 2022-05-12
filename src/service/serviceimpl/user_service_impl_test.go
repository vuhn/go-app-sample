package serviceimpl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vuhn/go-app-sample/errs"
	repoMocks "github.com/vuhn/go-app-sample/infrastructure/repository/mocks"
	pkgMocks "github.com/vuhn/go-app-sample/pkg/mocks"

	"github.com/vuhn/go-app-sample/testdata"
)

var ErrUnexpected = errors.New("Unexpected error")

type UserServiceTestSuite struct {
	suite.Suite
	userRepository *repoMocks.UserRepository
	token          *pkgMocks.Token
	password       *pkgMocks.Password
}

func (s *UserServiceTestSuite) SetupTest() {
	userRepositoryMock := new(repoMocks.UserRepository)
	tokenMock := new(pkgMocks.Token)
	passwordMock := new(pkgMocks.Password)
	s.userRepository = userRepositoryMock
	s.token = tokenMock
	s.password = passwordMock
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnSuccess() {
	user1 := testdata.User1
	s.userRepository.
		On("CreateUser", user1).
		Return(user1, nil)

	s.password.
		On("GenerateFromPassword", user1.Password).
		Return("encryptedPassword", nil)

	s.userRepository.
		On("GetUserByEmail", user1.Email).
		Return(nil, nil)

	userService := NewUserService(s.userRepository, s.token, s.password)
	createdUser, err := userService.CreateUser(user1)
	s.NoError(err)
	s.Equal(createdUser, user1)
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnErrorEmailExisted() {
	user2 := testdata.User2
	s.userRepository.
		On("GetUserByEmail", user2.Email).
		Return(user2, nil)

	userService := NewUserService(s.userRepository, s.token, s.password)
	createdUser, err := userService.CreateUser(user2)
	s.Equal(err, errs.ErrEmailExisted)
	s.Nil(createdUser)
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnServerError() {
	user2 := testdata.User2
	s.userRepository.
		On("GetUserByEmail", user2.Email).
		Return(nil, nil)

	s.password.
		On("GenerateFromPassword", user2.Password).
		Return("user2EncryptedPassword", nil)

	s.userRepository.
		On("CreateUser", user2).
		Return(nil, ErrUnexpected)

	userService := NewUserService(s.userRepository, s.token, s.password)

	createdUser, err := userService.CreateUser(user2)
	s.Equal(err, errs.ErrInternalServer)
	s.Nil(createdUser)
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnServerErrorWhenGenPassword() {
	user2 := testdata.User2
	s.userRepository.
		On("GetUserByEmail", user2.Email).
		Return(nil, nil)

	s.password.
		On("GenerateFromPassword", user2.Password).
		Return("", ErrUnexpected)

	userService := NewUserService(s.userRepository, s.token, s.password)

	createdUser, err := userService.CreateUser(user2)
	s.Equal(err, errs.ErrInternalServer)
	s.Nil(createdUser)
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnErrorWhenGetUserByEmail() {
	user2 := testdata.User2
	s.userRepository.
		On("GetUserByEmail", user2.Email).
		Return(nil, ErrUnexpected)

	s.userRepository.
		On("CreateUser", user2).
		Return(nil, nil)

	userService := NewUserService(s.userRepository, s.token, s.password)

	createdUser, err := userService.CreateUser(user2)
	s.Equal(err, errs.ErrInternalServer)
	s.Nil(createdUser)
}

func (s *UserServiceTestSuite) TestLogin_ShouldReturnSuccess() {
	user1 := testdata.User1

	email := user1.Email
	plainPassword := user1.Password
	encryptedPassword := "user1EncryptedPassword"

	// Encrypt user password
	user1.Password = encryptedPassword

	s.userRepository.
		On("GetUserByEmail", email).
		Return(user1, nil)

	s.password.
		On("CompareHashAndPassword", encryptedPassword, plainPassword).
		Return(true)

	claims := map[string]interface{}{
		"user_id": user1.ID,
		"email":   user1.Email,
	}

	generatedToken := "generated_token"
	s.token.
		On("Generate", claims).
		Return(generatedToken, nil)

	userService := NewUserService(s.userRepository, s.token, s.password)
	token, err := userService.Login(email, plainPassword)
	s.NoError(err)
	s.Equal(generatedToken, token)
}

func (s *UserServiceTestSuite) TestLogin_ShouldReturnEmailNotExists() {
	user1 := testdata.User1
	email := user1.Email
	plainPassword := user1.Password

	s.userRepository.
		On("GetUserByEmail", email).
		Return(nil, nil)

	userService := NewUserService(s.userRepository, s.token, s.password)
	token, err := userService.Login(email, plainPassword)
	s.Equal(errs.ErrEmailNotFound, err)
	s.Empty(token)
}

func (s *UserServiceTestSuite) TestLogin_ShouldReturnPasswordInvalid() {
	user1 := testdata.User1

	email := user1.Email
	plainPassword := "wrong_password"
	encryptedPassword := "user1EncryptedPassword"

	// Encrypt user password
	user1.Password = encryptedPassword

	s.userRepository.
		On("GetUserByEmail", email).
		Return(user1, nil)

	s.password.
		On("CompareHashAndPassword", encryptedPassword, plainPassword).
		Return(false)

	userService := NewUserService(s.userRepository, s.token, s.password)
	token, err := userService.Login(email, plainPassword)
	s.Equal(errs.ErrPasswordInvalid, err)
	s.Empty(token)
}

func (s *UserServiceTestSuite) TestLogin_ShouldReturnServerErrorWhenGetUserByEmail() {
	user1 := testdata.User1

	plainPassword := user1.Password
	email := user1.Email

	s.userRepository.
		On("GetUserByEmail", email).
		Return(nil, errs.ErrInternalServer)

	userService := NewUserService(s.userRepository, s.token, s.password)
	token, err := userService.Login(email, plainPassword)
	s.Equal(errs.ErrInternalServer, err)
	s.Empty(token)
}

func (s *UserServiceTestSuite) TestLogin_ShouldReturnServerErrorWhenGenJWTToken() {
	user1 := testdata.User1

	email := user1.Email
	plainPassword := user1.Password
	encryptedPassword := "user1EncryptedPassword"

	// Encrypt user password
	user1.Password = encryptedPassword

	s.userRepository.
		On("GetUserByEmail", email).
		Return(user1, nil)

	s.password.
		On("CompareHashAndPassword", encryptedPassword, plainPassword).
		Return(true)

	claims := map[string]interface{}{
		"user_id": user1.ID,
		"email":   user1.Email,
	}

	s.token.
		On("Generate", claims).
		Return("", ErrUnexpected)

	userService := NewUserService(s.userRepository, s.token, s.password)
	token, err := userService.Login(email, plainPassword)
	s.Equal(errs.ErrInternalServer, err)
	s.Empty(token)
}

func TestUserService(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserServiceTestSuite))
}
