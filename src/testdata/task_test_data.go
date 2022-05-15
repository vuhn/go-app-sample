package testdata

import (
	"time"

	"github.com/vuhn/go-app-sample/entity"
)

var Task1 = &entity.Task{
	ID:          "1",
	Title:       "Task 1",
	Description: "Task 1 description",
	Assignee:    "1",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

var Task2 = &entity.Task{
	ID:          "",
	Title:       "Task 1",
	Description: "Task 1 description",
	Assignee:    "1",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}
