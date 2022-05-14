package postgresrepo

import (
	"errors"

	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/repository"
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns impl of repository.UserRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func (u *postgresUserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *postgresUserRepository) GetUserByID(id string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.First(&user, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *postgresUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return user, nil
}
