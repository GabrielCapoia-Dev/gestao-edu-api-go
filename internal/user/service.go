package user

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{
		repository: repository,
	}
}