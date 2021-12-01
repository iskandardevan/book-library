package users

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/iskandardevan/book-library/app/middlewares"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx time.Duration
	JWTAuth *middlewares.ConfigJWT
}



func NewUseCase(userRepo UserRepoInterface, ctx time.Duration, JWTAuth *middlewares.ConfigJWT) *UserUseCase {
	return &UserUseCase{
		repo: 		userRepo,
		ctx:		ctx,
		JWTAuth: 	JWTAuth,
	}
}

func (usecase *UserUseCase) RegisterUser(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email belum di isi")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password belum di isi")
	}

	user, err := usecase.repo.RegisterUser(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	
	return user, nil
}

func (usecase *UserUseCase) LoginUser(email string, password string, ctx context.Context) (Domain, string, error){
	if email == "" {
		return Domain{},"", errors.New("email belum di isi")
	}
	if password == "" {
		return Domain{},"", errors.New("password belum di isi")
	
	}
	user, err := usecase.repo.GetEmail(ctx, email)
	if err != nil {
		return Domain{},"", err
	}
	token, errToken := usecase.JWTAuth.GenerateTokenJWT(user.Id)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", errors.New("token kosong bre")
	}
	return user, token, nil
}

func (usecase *UserUseCase) GetByID(id uint, ctx context.Context) (Domain, error){
	user, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada user dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return user, nil
}

func (usecase *UserUseCase) UpdateUserByID(id uint, ctx context.Context, domain Domain) (Domain, error){
	domain.Id = (id)
	user, err := usecase.repo.UpdateUserByID(id, ctx , domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada user dengan ID tersebut")
	}
	return user, nil
}


func (usecase *UserUseCase) DeleteUserByID(id uint, ctx context.Context) error{
	err := usecase.repo.DeleteUserByID(id, ctx)
	if err != nil {
		return errors.New("tidak ada user dengan ID tersebut")
	}
	return nil
}