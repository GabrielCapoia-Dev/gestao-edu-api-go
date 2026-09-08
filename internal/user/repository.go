package user

import (
	"context"
)

type Repository interface {
	EmailExists(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, user *User) error
}