package request

import "github.com/iskandardevan/book-library/business/users"

type RegisterUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (User *RegisterUser) ToDomain() *users.Domain {
	return & users.Domain{
		Email:User.Email  ,
		Name    :User.Name,
		Password :User.Password,
	}
}