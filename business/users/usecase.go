package users

import (
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx time.Duration
}



func NewUsecase(userRepo UserRepoInterface, ctx time.Duration) *UserUseCase {
	return &UserUseCase{
		repo: userRepo,
		ctx:	ctx,
	}
}

func (usecase *UserUseCase) RegisterUser(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email belum di isi")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password belum di isi")
	}
	data, err := usecase.repo.GetEmail(ctx, domain.Email)
	if err != nil {
		return Domain{}, err
	}
	if data.Id > 0 {
		return Domain{}, errors.New("email sudah dipakai")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password harus di isi") 
	}

	user, err := usecase.repo.RegisterUser(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	
	return user, nil
}

