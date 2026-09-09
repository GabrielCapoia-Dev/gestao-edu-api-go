package user

import (
	"api-golang/internal/security/password"
	"context"
	"errors"
	"strings"
)

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, request *CreateUserRequest) (*User, error) {

	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.ToLower(strings.TrimSpace(request.Email))

	if request.Name == "" {
		return nil, errors.New("Nome é obrigatório")
	}

	if request.Email == "" {
		return nil, errors.New("Email é obrigatório")
	}

	if request.Password == "" || len(request.Password) < 6 {
		return nil, errors.New("Password é obrigatório e deve ter pelo menos 6 caracteres")
	}

	if ok, _ := userService.repository.EmailExists(ctx, request.Email); ok {
		return nil, errors.New("Email ja cadastrado")
	}

	PasswordHash, err := password.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: PasswordHash,
	}

	if err := userService.repository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
