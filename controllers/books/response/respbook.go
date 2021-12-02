package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/books"
	"gorm.io/gorm"
)

type BookResponse struct {
	Id               uint           `json:"id"`
	Title            string         `json:"title"`
	Edition          int            `json:"edition"`
	Author_ID        uint           `json:"author_id"`
	Publisher_ID     uint           `json:"publisher_id"`
	Publication_Year string     `json:"publication_year"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
}

func FromDomainBook(domain books.Domain) BookResponse {
	return BookResponse{
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

func GetAllBook(data []books.Domain) []BookResponse{
	var res []BookResponse
	for _, val := range data{
		res = append(res, FromDomainBook(val))
	}
	return res
}