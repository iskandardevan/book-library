package books

import (
	"context"

	"github.com/iskandardevan/book-library/business/books"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *bookRepo {
	return &bookRepo{db: db}
}

func (Repo *bookRepo) RegisterBook(ctx context.Context, domain *books.Domain) (books.Domain, error) {
	book := Book{
		Id:               domain.Id,
		Title:            domain.Title,
		Edition:          domain.Edition,
		Author:           domain.Author,
		Publisher:        domain.Publisher,
		Publication_Year: domain.Publication_Year,
	}
	err := Repo.db.Create(&book)
	if err.Error != nil {
		return books.Domain{}, err.Error
	}
	return book.ToDomain(), nil
}
