package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)

type UserResponse struct {
	Id        	uint           `json:"id"`
	CreatedAt 	time.Time      `json:"createdAt"`
	UpdatedAt 	time.Time      `json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt `json:"deletedAt"`
	Email     	string         `json:"email"`
	Name      	string         `json:"name"`
	Age			int 			`json:"age"`
	Phone		string 			`json:"phone"`
	Password  	string         `json:"password"`
	Token     	string         `json:"token"`

}

type JWTResponse struct {
	Token string
	User interface{}
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Age: domain.Age,
		Phone: domain.Phone,
		Password:  domain.Password,
		Token:     domain.Token,
	}
}

func GetAllUsers(domain []users.Domain) []UserResponse {
	var GetAllUsers []UserResponse
	for _, val := range domain{
		GetAllUsers = append(GetAllUsers, FromDomain(val))
	}
	return GetAllUsers 
}