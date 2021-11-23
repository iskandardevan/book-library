package users

import (
	"time"
)

type User struct {
	Id 			uint 		`json:"id" form:"id"`
	Email		string 		`json:"email" form:"email"`
	FirstName 	string 		`json:"firstname" form:"firstname"`
	LastName 	string 		`json:"lastname" form:"lastname"`
	Age			int			`json:"age" form:"age"`
	Phone		string 		`json:"phonenum" form:"phonenum"`
	CreatedAt 	time.Time   `json:"createdAt"`
	UpdatedAt 	time.Time   `json:"updatedAt"`
}



//func (user User) ToDomain() users.Domain {
//	return users.Domain{
//
//
//	}
//}
//
//func FromDomain(domain users.Domain) User {
//	return User{
//
//
//	}
//}
