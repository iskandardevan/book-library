package response

import (
	"github.com/iskandardevan/book-library/business/users"
)


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