package dto

import (
	"time"

	"github.com/vuhn/go-app-sample/entity"
)

type (
	// CreateUserRequest defines http request params to create/update user
	CreateUserRequest struct {
		ID        string    `json:"id"`
		Email     string    `json:"email" validate:"required,email"`
		Fullname  string    `json:"full_name" validate:"required,gte=2,lte=100"`
		Password  string    `json:"password,omitempty" validate:"required,gte=6,lte=20"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// UserResponse defines http user response
	UserResponse struct {
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		Fullname  string    `json:"full_name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// UserLoginRequest defined http request params for user login
	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=6,lte=20"`
	}
)

// ToEntity returns user entity from user dto
func (u *CreateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Email:     u.Email,
		Fullname:  u.Fullname,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// NewUserResponseFromEntity returns UserResponse from user entity
func NewUserResponseFromEntity(entity *entity.User) *UserResponse {
	if entity == nil {
		return nil
	}
	return &UserResponse{
		ID:        entity.ID,
		Email:     entity.Email,
		Fullname:  entity.Fullname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
