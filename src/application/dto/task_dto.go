package dto

import (
	"time"

	"github.com/vuhn/go-app-sample/entity"
)

type (
	// CreateTaskRequest defines http request params to create/update user
	CreateTaskRequest struct {
		ID          string    `json:"-"`
		Title       string    `json:"title" validate:"required,lte=250"`
		Description string    `json:"description" validate:"lte=1000"`
		Assignee    string    `json:"assignee,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// TaskResponse defines http user response
	TaskResponse struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Assignee    string    `json:"assignee"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

// ToEntity returns user entity from user dto
func (t *CreateTaskRequest) ToEntity() *entity.Task {
	return &entity.Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Assignee:    t.Assignee,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// NewTaskResponseFromEntity returns TaskResponse from TaskEntity
func NewTaskResponseFromEntity(entity *entity.Task) *TaskResponse {
	return &TaskResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Assignee:    entity.Assignee,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
