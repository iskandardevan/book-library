package books

import (
	"context"
	"time"

	"github.com/iskandardevan/book-library/business/authors"
	"github.com/iskandardevan/book-library/business/publishers"
	"gorm.io/gorm"
)

type Domain struct {
	Id               	uint
	CreatedAt        	time.Time
	UpdatedAt        	time.Time
	DeletedAt        	gorm.DeletedAt 	`gorm:"index"`
	Title            	string
	Edition          	int
	Author_ID			uint
	Author				authors.Domain
	Publisher_ID		uint
	Publisher			publishers.Domain
	Publication_Year 	string
}

type BooksUsecaseInterface interface{
	AddBook(ctx context.Context, domain Domain) (Domain, error)
	GetAllBooks(ctx context.Context) ([]Domain, error)
}

type BooksRepoInterface interface {
	AddBook(ctx context.Context, domain Domain) (Domain, error)
	GetAllBooks(ctx context.Context) ([]Domain, error)
}
