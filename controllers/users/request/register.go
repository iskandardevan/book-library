package request

import "github.com/iskandardevan/book-library/business/users"

type RegisterUserRequest struct {
	Email    	string 	`json:"email"`
	Name     	string 	`json:"name"`
	Age			int 	`json:"age"`
	Phone		string 	`json:"phone"`
	Password 	string 	`json:"password"`
}

func (User *RegisterUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email	:User.Email  ,
		Name    :User.Name,
		Age		:User.Age,
		Phone	:User.Phone,
		Password:User.Password,
	}
}