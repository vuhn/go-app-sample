package entity

import "time"

// User defines table user structure
type User struct {
	ID        string
	Email     string
	Fullname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
