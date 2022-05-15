package entity

import (
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Assignee    string
	User        User `gorm:"foreignKey:Assignee"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
