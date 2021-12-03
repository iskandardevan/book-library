package authors

import (
	"time"

	"github.com/iskandardevan/book-library/business/authors"
	"gorm.io/gorm"
)


type Author struct {
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

func (author *Author) ToDomain() authors.Domain {
	res := authors.Domain{
		Id 			:author.Id,
		Email		:author.Email,
		Name 		:author.Name,
		Age			:author.Age,
		Phone		:author.Phone,
		Address		:author.Address,
		CreatedAt 	:author.CreatedAt,
		UpdatedAt 	:author.UpdatedAt,
	}
	// j, _ := json.MarshalIndent(res, " ", " ")
	// fmt.Println("To Domain Author", string(j))
	return res
}

func FromDomain(domain authors.Domain) Author  {
	return Author{
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

func GetAllAuthors(data []Author) []authors.Domain{
	res := []authors.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}