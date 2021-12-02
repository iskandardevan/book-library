package request

import "github.com/iskandardevan/book-library/business/reservations"

type AddReservationRequest struct {
	Book_ID uint `json:"book_id"`
	User_ID uint `json:"user_id"`
}

func (Reservation *AddReservationRequest) ToDomain() *reservations.Domain {
	return &reservations.Domain{
		Book_ID: Reservation.Book_ID,
		User_ID: Reservation.User_ID,
	}
}