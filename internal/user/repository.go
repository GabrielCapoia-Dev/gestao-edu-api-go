package user

import (
	"context"
)

type Repository interface {

	// Verifica se o email ja cadastrado e retorna dois valores, um booleano e um erro
	EmailExists(ctx context.Context, email string) (bool, error)


	Create(ctx context.Context, user *User) error
}