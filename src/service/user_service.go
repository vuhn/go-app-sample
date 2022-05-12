package service

import "github.com/vuhn/go-app-sample/entity"

// UserService defines methods for user service
type UserService interface {
	CreateUser(user *entity.User) (*entity.User, error)
	Login(email string, password string) (string, error)
}
