package users

import (
	"context"

	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (Repo *userRepo) RegisterUser(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	user := User{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Age			:domain.Age,
		Phone		:domain.Phone,
	}
	err := Repo.db.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
} 

