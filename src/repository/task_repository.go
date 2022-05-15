package repository

import "github.com/vuhn/go-app-sample/entity"

type TaskRepository interface {
	CreateTask(task *entity.Task) (*entity.Task, error)
}
