package serviceimpl

import (
	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/errs"
	"github.com/vuhn/go-app-sample/repository"
	"github.com/vuhn/go-app-sample/service"
)

// NewTaskService return implementation of task service interface
func NewTaskService(
	userRepository repository.UserRepository,
	taskRepository repository.TaskRepository,
) service.TaskService {
	return &TaskServiceImpl{
		userRepository: userRepository,
		taskRepository: taskRepository,
	}
}

// TaskServiceImpl implements TaskService interface
type TaskServiceImpl struct {
	userRepository repository.UserRepository
	taskRepository repository.TaskRepository
}

// CreateTask is method to create a Task
func (t *TaskServiceImpl) CreateTask(task *entity.Task) (*entity.Task, error) {
	if task.Assignee != "" {
		user, err := t.userRepository.GetUserByID(task.Assignee)
		if err != nil {
			return nil, errs.ErrInternalServer
		}

		if user == nil {
			return nil, errs.ErrUserNotFound
		}
	}

	taskEntity, err := t.taskRepository.CreateTask(task)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	return taskEntity, nil
}
