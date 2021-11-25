package publishers

import (
	"time"

	"github.com/iskandardevan/book-library/business/publishers"
	"gorm.io/gorm"
)


type Publisher struct {
	Id 			uint 			`gorm:"primaryKey"`
	Email		string			`gorm:"unique"`
	Name 		string 		
	Age			int			
	Phone		string 
	Address		string
	CreatedAt 	time.Time   
	UpdatedAt 	time.Time   
	DeletedAt	gorm.DeletedAt 	`gorm:"index"`
}

func (publisher *Publisher) ToDomain() publishers.Domain {
	return publishers.Domain{
		Id 			:publisher.Id,
		Email		:publisher.Email,
		Name 		:publisher.Name,
		Age			:publisher.Age,
		Phone		:publisher.Phone,
		Address		:publisher.Address,
		CreatedAt 	:publisher.CreatedAt,
		UpdatedAt 	:publisher.UpdatedAt,
	}
}

func FromDomain(domain publishers.Domain) Publisher  {
	return Publisher{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Age			:domain.Age,
		Phone		:domain.Phone,
		Address		:domain.Address,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}