package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/publishers"
	"gorm.io/gorm"
)

type PublisherResponse struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
}

func FromDomainPublisher(domain publishers.Domain) PublisherResponse {
	return PublisherResponse{
		Id:        domain.Id,
		Email:     domain.Email,
		Name:      domain.Name,
		Age:       domain.Age,
		Phone:     domain.Phone,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}