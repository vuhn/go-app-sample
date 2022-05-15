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

func (u *postgresTaskRepository) CreateTask(task *entity.Task) (*entity.Task, error) {
	if err := u.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
