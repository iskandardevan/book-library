package request

import "github.com/iskandardevan/book-library/business/publishers"

type AddPublisherRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (Publisher *AddPublisherRequest) ToDomain() *publishers.Domain {
	return &publishers.Domain{
		Email:   Publisher.Email,
		Name:    Publisher.Name,
		Age:     Publisher.Age,
		Phone:   Publisher.Phone,
		Address: Publisher.Address,
	}
}