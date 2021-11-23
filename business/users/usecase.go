package users

import (
	"github.com/iskandardevan/book-library/controllers/users/request"
)

type UserUseCase struct {
	repo UserRepoInterface
}

func NewUsecase(userRepo UserRepoInterface) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
	}
}

func (usecase *UserUseCase) CreateUser(user request.CreateUser ) ([]Domain, error) {
	return []Domain{}, nil
}
