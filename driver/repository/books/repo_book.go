package books

import (
	"context"

	"github.com/iskandardevan/book-library/business/books"
	"gorm.io/gorm"
)

type bookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(DB *gorm.DB) *bookRepo {
	return &bookRepo{DB: DB}
}

func (Repo *bookRepo) RegisterBook(ctx context.Context, domain *books.Domain) (books.Domain, error) {
	book := Book{
		Id:               domain.Id,
		Title:            domain.Title,
		Edition:          domain.Edition,
		Author_ID			:domain.Author_ID,
		Publisher_ID		:domain.Publisher_ID,
		Publication_Year: domain.Publication_Year,
	}
	err := Repo.DB.Create(&book)
	if err.Error != nil {
		return books.Domain{}, err.Error
	}
	return book.ToDomain(), nil
}

func (Repo *bookRepo) GetAllBooks(ctx context.Context) ([]books.Domain, error){
	var book []Book
	err := Repo.DB.Find(&book)
	if err.Error != nil {
		return []books.Domain{}, err.Error
	}
	return GetAllBook(book), nil
}