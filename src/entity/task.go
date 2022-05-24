package entity

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Assignee    sql.NullString
	User        *User `gorm:"foreignKey:Assignee"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
