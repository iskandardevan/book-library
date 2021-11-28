package request

import "github.com/iskandardevan/book-library/business/authors"

type AddAuthorRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (Author *AddAuthorRequest) ToDomain() *authors.Domain {
	return &authors.Domain{
		Email	: Author.Email,
		Name	: Author.Name,
		Age		: Author.Age,
		Phone	: Author.Phone,
		Address	: Author.Address,
	}
}