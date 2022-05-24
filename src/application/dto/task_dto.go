package dto

import (
	"database/sql"
	"time"

	"github.com/vuhn/go-app-sample/entity"
)

type (
	// CreateTaskRequest defines http request params to create/update tasks
	CreateTaskRequest struct {
		ID          string    `json:"-"`
		Title       string    `json:"title" validate:"required,lte=250"`
		Description string    `json:"description" validate:"lte=1000"`
		Assignee    string    `json:"assignee,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// GetTasksListRequest defines http request params to get tasks list
	GetTasksListRequest struct {
		Limit  int `query:"limit" validate:"required,gte=0,lte=1000"`
		Offset int `query:"offset" validate:"gte=0"`
	}

	// TaskResponse defines http user response
	TaskResponse struct {
		ID          string        `json:"id"`
		Title       string        `json:"title"`
		Description string        `json:"description"`
		Assignee    *UserResponse `json:"assignee"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
	}
)

// ToEntity returns user entity from user dto
func (t *CreateTaskRequest) ToEntity() *entity.Task {
	isAssigneeNull := false
	if t.Assignee != "" {
		isAssigneeNull = true
	}
	return &entity.Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Assignee:    sql.NullString{Valid: isAssigneeNull, String: t.Assignee},
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// NewTaskResponseFromEntity returns TaskResponse from TaskEntity
func NewTaskResponseFromEntity(entity *entity.Task) *TaskResponse {
	if entity == nil {
		return nil
	}
	return &TaskResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Assignee:    NewUserResponseFromEntity(entity.User),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// NewTasksResponseFromEntities returns TaskResponses List from Task Entities
func NewTasksResponseFromEntities(entities []*entity.Task) []*TaskResponse {
	tasksResponse := []*TaskResponse{}
	for _, entity := range entities {
		task := NewTaskResponseFromEntity(entity)
		tasksResponse = append(tasksResponse, task)
	}
	return tasksResponse
}
