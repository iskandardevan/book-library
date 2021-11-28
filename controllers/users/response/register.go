package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)

type UserRegister struct {
	Id        	uint           	`json:"id"`
	CreatedAt 	time.Time      	`json:"createdAt"`
	UpdatedAt 	time.Time      	`json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt 	`json:"deletedAt"`
	Email     	string         	`json:"email"`
	Name      	string         	`json:"name"`
	Age			int 			`json:"age"`
	Phone		string 			`json:"phone"`
	Password  	string         	`json:"password"`
	Token     	string         	`json:"token"`
}

func FromUserRegister(domain users.Domain) UserResponse {
	return UserResponse{
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		DeletedAt	:domain.DeletedAt,
		Name		:domain.Name,
		Email		:domain.Email,
		Age			:domain.Age,
		Phone		:domain.Phone,
		Password	:domain.Password,
		Token		:domain.Token,
	}
}