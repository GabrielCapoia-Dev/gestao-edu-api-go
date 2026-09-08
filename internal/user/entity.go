package user

import "time"

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
