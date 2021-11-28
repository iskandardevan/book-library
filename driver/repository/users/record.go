package users

import (
	"time"

	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)


type User struct {
	Id 			uint 			`gorm:"primaryKey"`
	Email		string			`gorm:"unique"`
	Name 		string 		
	Age			int			
	Phone		string
	Password	string
	Token		string
	CreatedAt 	time.Time   
	UpdatedAt 	time.Time   
	DeletedAt	gorm.DeletedAt 	`gorm:"index"`
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id 			:user.Id,
		Email		:user.Email,
		Name 		:user.Name,
		Age			:user.Age,
		Phone		:user.Phone,
		Password	:user.Password,
		Token		:user.Token,
		CreatedAt 	:user.CreatedAt,
		UpdatedAt 	:user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User  {
	return User{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Age			:domain.Age,
		Phone		:domain.Phone,
		Password	:domain.Password,
		Token		:domain.Token,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}