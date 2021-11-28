package authors

import (
	"context"

	"github.com/iskandardevan/book-library/business/authors"
	"gorm.io/gorm"
)

type authorRepo struct {
	DB *gorm.DB
}

func NewAuthorRepo(DB *gorm.DB) *authorRepo {
	return &authorRepo{DB: DB}
}

func (Repo *authorRepo) RegisterAuthor(ctx context.Context, domain *authors.Domain) (authors.Domain, error) {
	author := Author{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Age			:domain.Age,
		Phone		:domain.Phone,
		Address		:domain.Address,
	}
	err := Repo.DB.Create(&author)
	if err.Error != nil {
		return authors.Domain{}, err.Error
	}
	return author.ToDomain(), nil
} 