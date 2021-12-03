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
	Author			 authors.Author				`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Publisher_ID	 uint						
	Publisher		 publishers.Publisher		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Publication_Year string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt 			`gorm:"index"`
}

func (book *Book) ToDomain() books.Domain {
	res := books.Domain{
		Id					:book.Id,
		Title            	:book.Title,
		Edition          	:book.Edition,
		Author_ID			:book.Author_ID,
		Author				:book.Author.ToDomain() ,
		Publisher_ID		:book.Publisher_ID,
		Publisher			:book.Publisher.ToDomain() ,
		Publication_Year 	:book.Publication_Year,
		CreatedAt			:book.CreatedAt,
		UpdatedAt			:book.UpdatedAt,
	}
	// j, _ := json.MarshalIndent(res, " ", " ")
	// fmt.Println("To Domain Book", string(j))
	return res
}

func FromDomain(domain books.Domain) Book {
	return Book{
		Id					:domain.Id,
		Title            	:domain.Title,
		Edition          	:domain.Edition,
		Author_ID			:domain.Author_ID,
		Author				:authors.FromDomain(domain.Author) ,
		Publisher_ID		:domain.Publisher_ID,
		Publisher			:publishers.FromDomain(domain.Publisher) ,
		Publication_Year 	:domain.Publication_Year,
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}

}

func GetAllBook(data []Book) []books.Domain{
	res := []books.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	// j, _ := json.MarshalIndent(res, " ", " ")
	// fmt.Println("isi GetAllBook" , string(j))
	// fmt.Println("==============================")
	return res
}