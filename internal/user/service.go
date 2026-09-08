package user

import "context"

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, name, email, password string) (*User, error) {

	// Lógica para criar um usuário, incluindo validação de email, hash da senha, etc.
	return nil, nil
}
