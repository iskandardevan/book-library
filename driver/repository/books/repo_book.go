package books

import (
	"context"
	"errors"

	"github.com/iskandardevan/book-library/business/books"
	"gorm.io/gorm"
)

type bookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(DB *gorm.DB) *bookRepo {
	return &bookRepo{DB: DB}
}

func (Repo *bookRepo) AddBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
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
	err := Repo.DB.Preload("Publisher").Preload("Author").Find(&book)
	if err.Error != nil {
		return []books.Domain{}, err.Error
	}
	// j, _ := json.MarshalIndent(book, " ", " ")
	// fmt.Println("isi book", string(j))
	// fmt.Println("=======================================")
	
	return GetAllBook(book), nil
	
}

func (Repo *bookRepo) Delete(id uint, ctx context.Context) error{
	book := Book{}
	err := Repo.DB.Delete(&book, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}