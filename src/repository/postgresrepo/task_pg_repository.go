package postgresrepo

import (
	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/repository"
	"gorm.io/gorm"
)

type postgresTaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository returns impl of TaskRepository
func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &postgresTaskRepository{
		db: db,
	}
}

func (t *postgresTaskRepository) CreateTask(task *entity.Task) (*entity.Task, error) {
	if err := t.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *postgresTaskRepository) GetTasksList(limit int, offset int) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	err := t.db.Debug().
		Joins("User").
		Limit(limit).
		Offset(offset).
		Find(&tasks).Error

	if err != nil {
		return tasks, 0, err
	}

	// Count all tasks
	var total int64
	if err := t.db.Model(&entity.Task{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}
