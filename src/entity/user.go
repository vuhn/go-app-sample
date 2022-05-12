package entity

import "time"

// User reflects users data from DB
type User struct {
	ID        string
	Email     string
	Fullname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
