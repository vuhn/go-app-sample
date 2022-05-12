package testdata

import (
	"time"

	"github.com/google/uuid"
	"github.com/vuhn/go-app-sample/entity"
)

var now = time.Now()
var User1 = &entity.User{
	ID:        uuid.NewString(),
	Email:     "test1@example.com",
	Fullname:  "Test1",
	Password:  "123456",
	CreatedAt: now,
	UpdatedAt: now,
}

var User2 = &entity.User{
	ID:        uuid.NewString(),
	Email:     "test2@example.com",
	Fullname:  "Test2",
	Password:  "123456",
	CreatedAt: now,
	UpdatedAt: now,
}

var User3 = &entity.User{
	ID:        "",
	Email:     "test1@example.com",
	Fullname:  "Test1",
	Password:  "123456",
	CreatedAt: now,
	UpdatedAt: now,
}
