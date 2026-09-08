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

func (userService *UserService) CreateUser(ctx context.Context, request *CreateUserRequest) (*User, error) {

	// Lógica para criar um usuário, incluindo validação de email, hash da senha, etc.
	return nil, nil
}
