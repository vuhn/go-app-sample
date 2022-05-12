package serviceimpl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vuhn/go-app-sample/errs"
	"github.com/vuhn/go-app-sample/infrastructure/repository/mocks"
	"github.com/vuhn/go-app-sample/testdata"
)

var ErrUnexpected = errors.New("Unexpected error")

type UserServiceTestSuite struct {
	suite.Suite
	userRepository *mocks.UserRepository
}

func (s *UserServiceTestSuite) SetupTest() {
	userRepositoryMock := new(mocks.UserRepository)
	s.userRepository = userRepositoryMock
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnSuccess() {
	user1 := testdata.User1
	s.userRepository.
		On("CreateUser", user1).
		Return(user1, nil)

	s.userRepository.
		On("GetUserByEmail", user1.Email).
		Return(nil, nil)

	userService := NewUserService(s.userRepository)
	createdUser, err := userService.CreateUser(user1)
	s.NoError(err)
	s.Equal(createdUser, user1)
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnErrorEmailExisted() {
	user2 := testdata.User2
	s.userRepository.
		On("GetUserByEmail", user2.Email).
		Return(user2, nil)

	userService := NewUserService(s.userRepository)
	createdUser, err := userService.CreateUser(user2)
	s.Equal(err, errs.ErrEmailExisted)
	s.Nil(createdUser)
}

func (s *UserServiceTestSuite) TestCreateUser_ShouldReturnServerError() {
	user2 := testdata.User2
	s.userRepository.
		On("GetUserByEmail", user2.Email).
		Return(nil, nil)

	s.userRepository.
		On("CreateUser", user2).
		Return(nil, ErrUnexpected)

	userService := NewUserService(s.userRepository)

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

	userService := NewUserService(s.userRepository)

	createdUser, err := userService.CreateUser(user2)
	s.Equal(err, errs.ErrInternalServer)
	s.Nil(createdUser)
}

func TestUserService(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserServiceTestSuite))
}
