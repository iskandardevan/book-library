package books

import (
	"time"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/iskandardevan/book-library/driver/repository/authors"
	"github.com/iskandardevan/book-library/driver/repository/publishers"
	"gorm.io/gorm"
)

type Book struct {
	Id               uint   					`gorm:"primaryKey"`
	Title          	 string 					`gorm:"unique"`
	Edition          int
	Author_ID		 uint			
	Author			 authors.Author				`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	Publisher_ID	 uint						`gorm:"foreignKey"`
	Publisher		 publishers.Publisher		`gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	Publication_Year string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt 			`gorm:"index"`
}

func (book *Book) ToDomain() books.Domain {
	return books.Domain{
		Id					:book.Id,
		Title            	:book.Title,
		Edition          	:book.Edition,
		Author_ID			:book.Author_ID,
		Publisher_ID		:book.Publisher_ID,
		Publication_Year 	:book.Publication_Year,
		CreatedAt			:book.CreatedAt,
		UpdatedAt			:book.UpdatedAt,
	}
}

func FromDomain(domain books.Domain) Book {
	return Book{
		Id					:domain.Id,
		Title				:domain.Title,
		Edition				:domain.Edition,
		Author_ID			:domain.Author_ID,
		Publisher_ID		:domain.Publisher_ID,
		Publication_Year	:domain.Publication_Year,
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}

}