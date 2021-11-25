package books

import (
	"time"

	"github.com/iskandardevan/book-library/business/books"
	"gorm.io/gorm"
)

type Book struct {
	Id               uint   `gorm:"primaryKey"`
	Title            string `gorm:"unique"`
	Edition          int
	Author           string
	Publisher        string
	Publication_Year time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (book *Book) ToDomain() books.Domain {
	return books.Domain{
		Id:        book.Id,
		Title            	: book.Title,
		Edition          	:book.Edition,
		Author				:book.Author,
		Publisher			:book.Publisher,
		Publication_Year 	:book.Publication_Year,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func FromDomain(domain books.Domain) Book {
	return Book{
		Id:        			domain.Id,
		Title:  			domain.Title,
		Edition:			domain.Edition,
		Author:				domain.Author,
		Publisher:			domain.Publisher,
		Publication_Year:	domain.Publication_Year,
		CreatedAt:			domain.CreatedAt,
		UpdatedAt: 			domain.UpdatedAt,
	}

}