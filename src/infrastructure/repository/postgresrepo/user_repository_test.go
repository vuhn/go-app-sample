package postgresrepo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/vuhn/go-app-sample/infrastructure/repository"
	"github.com/vuhn/go-app-sample/testdata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite
	mockSQL    sqlmock.Sqlmock
	repository repository.UserRepository
}

func (s *UserRepoTestSuite) SetupTest() {
	sqlDB, mockSQL, err := sqlmock.New()
	s.NoError(err)

	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	s.mockSQL = mockSQL

	repo := NewUserRepository(mockDB)
	s.repository = repo
}

func (s *UserRepoTestSuite) TestCreateUser_ShouldReturnSuccess() {
	user := testdata.User1
	result := sqlmock.NewResult(0, 1)
	sqlCreate := `INSERT INTO "users" ("id","email","fullname","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`

	s.mockSQL.ExpectBegin()
	s.mockSQL.ExpectExec(regexp.QuoteMeta(sqlCreate)).
		WithArgs(user.ID, user.Email, user.Fullname, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnResult(result)

	s.mockSQL.ExpectCommit()

	createdUser, err := s.repository.CreateUser(user)

	s.Equal(err, nil)
	s.Equal(createdUser, user)
}
func (s *UserRepoTestSuite) TestCreateUser_ShouldReturnError() {
	sqlCreate := `INSERT INTO "users" ("id","email","fullname","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`

	user := testdata.User3
	s.mockSQL.ExpectBegin()
	s.mockSQL.ExpectExec(regexp.QuoteMeta(sqlCreate)).
		WithArgs(user.ID, user.Email, user.Fullname, user.Password, user.CreatedAt, user.UpdatedAt).
		WillReturnError(gorm.ErrPrimaryKeyRequired)

	s.mockSQL.ExpectRollback()

	createdUser, err := s.repository.CreateUser(user)

	s.EqualError(err, gorm.ErrPrimaryKeyRequired.Error())
	s.Nil(createdUser)
}

func (s *UserRepoTestSuite) TestGetUserById_ShouldReturnUser() {
	user := testdata.User1
	result := sqlmock.NewRows([]string{"id", "email", "fullname", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Email, user.Fullname, user.Password, user.CreatedAt, user.UpdatedAt)
	sqlSelect := `SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`

	s.mockSQL.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
		WithArgs(user.ID).
		WillReturnRows(result)

	returnUser, err := s.repository.GetUserByID(user.ID)

	s.NoError(err)
	s.Equal(returnUser, user)
}

func (s *UserRepoTestSuite) TestGetUserById_ShouldReturnNilWithNoError() {
	user := testdata.User1
	sqlSelect := `SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`

	s.mockSQL.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
		WithArgs(user.ID).
		WillReturnError(gorm.ErrRecordNotFound)

	returnUser, err := s.repository.GetUserByID(user.ID)

	s.NoError(err)
	s.Nil(returnUser)
}
func (s *UserRepoTestSuite) TestGetUserById_ShouldReturnError() {
	user := testdata.User1
	sqlSelect := `SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`

	s.mockSQL.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
		WithArgs(user.ID).
		WillReturnError(gorm.ErrInvalidData)

	returnUser, err := s.repository.GetUserByID(user.ID)

	s.EqualError(err, gorm.ErrInvalidData.Error())
	s.Nil(returnUser)
}

func (s *UserRepoTestSuite) TestGetUserByEmail_SouldReturnUserByGivenEmail() {
	user := testdata.User1
	result := sqlmock.NewRows([]string{"id", "email", "fullname", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Email, user.Fullname, user.Password, user.CreatedAt, user.UpdatedAt)
	sqlSelect := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`

	s.mockSQL.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
		WithArgs(user.Email).
		WillReturnRows(result)

	returnUser, err := s.repository.GetUserByEmail(user.Email)

	s.NoError(err)
	s.Equal(returnUser, user)
}

func (s *UserRepoTestSuite) TestGetUserByEmail_SouldReturnNilWithNoError() {
	user := testdata.User1
	sqlSelect := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`

	s.mockSQL.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
		WithArgs(user.Email).
		WillReturnError(gorm.ErrRecordNotFound)

	returnUser, err := s.repository.GetUserByEmail(user.Email)

	s.NoError(err)
	s.Nil(returnUser)
}

func (s *UserRepoTestSuite) TestGetUserByEmail_SouldReturnError() {
	user := testdata.User1
	sqlSelect := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`

	s.mockSQL.ExpectQuery(regexp.QuoteMeta(sqlSelect)).
		WithArgs(user.Email).
		WillReturnError(gorm.ErrInvalidData)

	returnUser, err := s.repository.GetUserByEmail(user.Email)

	s.EqualError(err, gorm.ErrInvalidData.Error())
	s.Nil(returnUser)
}

func TestPostgresUserRepository(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserRepoTestSuite))
}
