package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/authors"
	"gorm.io/gorm"
)

type AuthorResponse struct {
	Id       	uint         	`json:"id"`
	CreatedAt 	time.Time    	`json:"createdAt"`
	UpdatedAt 	time.Time    	`json:"updatedAt"`
	DeletedAt	gorm.DeletedAt 	`json:"deletedAt"`
	Email    	string       	`json:"email"`
	Name     	string       	`json:"name"`
	Age			int 			`json:"age"`
	Phone		string 			`json:"phone"`
	Address		string			`json:"address"`

}

func FromDomainAuthor(domain authors.Domain) AuthorResponse {
	return AuthorResponse{
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

// func GetAllauthor(domain []author.Domain) []UserResponse {
// 	var GetAllauthor []UserResponse
// 	for _, val := range domain{
// 		GetAllauthor = append(GetAllauthor, FromDomain(val))
// 	}
// 	return GetAllauthor 
// }