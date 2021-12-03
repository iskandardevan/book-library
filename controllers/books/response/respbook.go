package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/books"
	respAuth "github.com/iskandardevan/book-library/controllers/authors/response"
	respPub "github.com/iskandardevan/book-library/controllers/publishers/response"
	"gorm.io/gorm"
)

type BookResponse struct {
	Id               uint           `json:"id"`
	Title            string         `json:"title"`
	Edition          int            `json:"edition"`
	Author_ID        uint           `json:"author_id"`
	Author			 respAuth.AuthorResponse
	Publisher_ID     uint           `json:"publisher_id"`
	Publisher		 respPub.PublisherResponse
	Publication_Year string     	`json:"publication_year"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
}

func FromDomainBook(domain books.Domain) BookResponse {
	return BookResponse{
		Id					:domain.Id,
		Title            	:domain.Title,
		Edition          	:domain.Edition,
		Author_ID			:domain.Author_ID,
		Author				:respAuth.FromDomainAuthor(domain.Author) ,
		Publisher_ID		:domain.Publisher_ID,
		Publisher			:respPub.FromDomainPublisher(domain.Publisher) ,
		Publication_Year 	:domain.Publication_Year,
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}
}

func GetAllBook(data []books.Domain) []BookResponse{
	var res []BookResponse
	for _, val := range data{
		res = append(res, FromDomainBook(val))
	}
	return res
}