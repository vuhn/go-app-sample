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
	if task.Assignee.String != "" {
		user, err := t.userRepository.GetUserByID(task.Assignee.String)
		if err != nil {
			return nil, errs.ErrInternalServer
		}

		if user == nil {
			return nil, errs.ErrUserNotFound
		}
		task.User = user
	}

	taskEntity, err := t.taskRepository.CreateTask(task)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	return taskEntity, nil
}

// GetTasksList return lists of tasks
func (t *TaskServiceImpl) GetTasksList(limit int, offset int) ([]*entity.Task, int64, error) {
	tasks, total, err := t.taskRepository.GetTasksList(limit, offset)
	if err != nil {
		return nil, 0, errs.ErrInternalServer
	}

	return tasks, total, nil
}
