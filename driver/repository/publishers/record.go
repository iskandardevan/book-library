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
	Phone		string 
	Address		string
	CreatedAt 	time.Time   
	UpdatedAt 	time.Time   
	DeletedAt	gorm.DeletedAt 	`gorm:"index"`
}

func (publisher *Publisher) ToDomain() publishers.Domain {
	res := publishers.Domain{
		Id 			:publisher.Id,
		Email		:publisher.Email,
		Name 		:publisher.Name,
		Phone		:publisher.Phone,
		Address		:publisher.Address,
		CreatedAt 	:publisher.CreatedAt,
		UpdatedAt 	:publisher.UpdatedAt,
	}
	// j, _ := json.MarshalIndent(res, " ", " ")
	// fmt.Println("To Domain Publiser", string(j))
	return res
}

func FromDomain(domain publishers.Domain) Publisher  {
	return Publisher{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Phone		:domain.Phone,
		Address		:domain.Address,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}

func GetAllPublishers(data []Publisher) []publishers.Domain{
	res := []publishers.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}