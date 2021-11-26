package users

import (
	"context"

	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *userRepo {
	return &userRepo{DB: DB,}
}

func (Repo *userRepo) RegisterUser(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	user := FromDomain(*domain)
	err := Repo.DB.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
} 

func (Repo *userRepo) GetEmail(ctx context.Context, email string) (users.Domain, error){
	var user User
	err := Repo.DB.Find(&user, "email=?", email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}
