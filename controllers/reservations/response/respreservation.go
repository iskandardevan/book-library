package response

import (
	"time"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/iskandardevan/book-library/business/reservations"
	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)

type ReservationResponse struct {
	Id        uint           `json:"id"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
	Book      books.Domain   `json:"book"`
	Book_ID   uint           `json:"book_id"`
	User      users.Domain   `json:"user"`
	User_ID   uint           `json:"user_id"`
	Start     time.Time      `json:"start"`
	End       time.Time      `json:"end"`
}

func FromDomainReservation(domain reservations.Domain) ReservationResponse{
	return ReservationResponse {
		Id:    				domain.Id,
		Book_ID:  			domain.Book_ID,
		User_ID:  			domain.User_ID,
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