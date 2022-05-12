package dto

import (
	"time"

	"github.com/vuhn/go-app-sample/entity"
)

type (
	// UserRequest defines http request params
	UserRequest struct {
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
		Email     string    `json:"email" validate:"required,email"`
		Fullname  string    `json:"full_name" validate:"required,gte=2,lte=100"`
		Password  string    `json:"password,omitempty" validate:"required,gte=6,lte=20"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

// ToEntity returns user entity from user dto
func (u *UserRequest) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Email:     u.Email,
		Fullname:  u.Fullname,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// NewUserResponseFromEntity returns UserResponse from user entity
func NewUserResponseFromEntity(entity *entity.User) *UserResponse {
	return &UserResponse{
		ID:        entity.ID,
		Email:     entity.Email,
		Fullname:  entity.Fullname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
