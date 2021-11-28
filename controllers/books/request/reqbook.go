package request

import (
	"time"

	"github.com/iskandardevan/book-library/business/books"
)

type AddBookRequest struct {
	Title            string    `json:"title"`
	Edition          int       `json:"edition"`
	Author_ID        uint      `json:"author_id"`
	Publisher_ID     uint      `json:"publisher_id"`
	Publication_Year time.Time `json:"publication_year"`
}

func (Book *AddBookRequest) ToDomain() *books.Domain {
	return &books.Domain{
		Title:            Book.Title,
		Edition:          Book.Edition,
		Author_ID:        Book.Author_ID,
		Publisher_ID:     Book.Publisher_ID,
		Publication_Year: Book.Publication_Year,
	}
}