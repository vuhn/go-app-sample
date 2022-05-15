package postgresrepo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/vuhn/go-app-sample/repository"
	"github.com/vuhn/go-app-sample/testdata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskRepoTestSuite struct {
	suite.Suite
	mockSQL    sqlmock.Sqlmock
	repository repository.TaskRepository
}

func (s *TaskRepoTestSuite) SetupTest() {
	sqlDB, mockSQL, err := sqlmock.New()
	s.NoError(err)

	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	s.mockSQL = mockSQL

	repo := NewTaskRepository(mockDB)
	s.repository = repo
}

func (s *TaskRepoTestSuite) TestCreateTask_ShouldReturnSuccess() {
	task := testdata.Task1
	result := sqlmock.NewResult(0, 1)
	sqlCreate := `INSERT INTO "tasks" ("id","title","description","assignee","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`

	s.mockSQL.ExpectBegin()
	s.mockSQL.ExpectExec(regexp.QuoteMeta(sqlCreate)).
		WithArgs(task.ID, task.Title, task.Description, task.Assignee, task.CreatedAt, task.UpdatedAt).
		WillReturnResult(result)

	s.mockSQL.ExpectCommit()

	createdTask, err := s.repository.CreateTask(task)

	s.Equal(err, nil)
	s.Equal(createdTask, task)
}
func (s *TaskRepoTestSuite) TestCreateTask_ShouldReturnError() {
	sqlCreate := `INSERT INTO "tasks" ("id","title","description","assignee","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`

	task := testdata.Task2
	s.mockSQL.ExpectBegin()
	s.mockSQL.ExpectExec(regexp.QuoteMeta(sqlCreate)).
		WithArgs(task.ID, task.Title, task.Description, task.Assignee, task.CreatedAt, task.UpdatedAt).
		WillReturnError(gorm.ErrPrimaryKeyRequired)

	s.mockSQL.ExpectRollback()

	createdTask, err := s.repository.CreateTask(task)

	s.EqualError(err, gorm.ErrPrimaryKeyRequired.Error())
	s.Nil(createdTask)
}

func TestPostgresTaskRepository(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TaskRepoTestSuite))
}
