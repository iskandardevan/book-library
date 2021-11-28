package request

import "github.com/iskandardevan/book-library/business/users"

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(login UserLoginRequest) users.Domain {
	return users.Domain{
		Email: login.Email,
		Password: login.Password,
	}
}