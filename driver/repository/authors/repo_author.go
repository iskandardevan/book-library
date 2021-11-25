package authors

import (
	"context"

	"github.com/iskandardevan/book-library/business/authors"
	"gorm.io/gorm"
)

type authorRepo struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *authorRepo {
	return &authorRepo{db: db}
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
	err := Repo.db.Create(&author)
	if err.Error != nil {
		return authors.Domain{}, err.Error
	}
	return author.ToDomain(), nil
} 