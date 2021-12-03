package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/reservations"
	respBook "github.com/iskandardevan/book-library/controllers/books/response"
	respUser "github.com/iskandardevan/book-library/controllers/users/response"
	"gorm.io/gorm"
)

type ReservationResponse struct {
	Id        		 uint           `json:"id"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
	Book      		 respBook.BookResponse
	Book_ID   		 uint           `json:"book_id"`
	User      		 respUser.UserResponse
	User_ID   		 uint           `json:"user_id"`
	Start    		 time.Time      `json:"start"`
	End      		 time.Time      `json:"end"`
}

func FromDomainReservation(domain reservations.Domain) ReservationResponse{
	return ReservationResponse {
		Id:    				domain.Id,
		Book_ID:  			domain.Book_ID,
		Book: 				respBook.FromDomainBook(domain.Book),
		User_ID:  			domain.User_ID,
		User:				respUser.FromDomain(domain.User),
		Start:				domain.Start,
		End:   				domain.End,
		CreatedAt:			domain.CreatedAt,
		UpdatedAt:			domain.UpdatedAt,
	}
}

func GetAllReservation(data []reservations.Domain) []ReservationResponse{
	var res []ReservationResponse
	for _, val := range data{
		res = append(res, FromDomainReservation(val))
	}
	return res
}