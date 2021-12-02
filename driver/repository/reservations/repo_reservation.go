package reservations

import (
	"context"

	"github.com/iskandardevan/book-library/business/reservations"
	"gorm.io/gorm"
)

type reservationRepo struct {
	DB *gorm.DB
}

func NewReservationRepo(DB *gorm.DB) *reservationRepo {
	return &reservationRepo{DB: DB}
}

func (Repo *reservationRepo) AddReservation(ctx context.Context, domain reservations.Domain) (reservations.Domain, error) {
	reservation := Reservation{
		Id:    				domain.Id,
		Book_ID:  			domain.Book_ID,
		User_ID:  			domain.User_ID,
		Start:				domain.Start,
		End:   				domain.End,
	}
	err := Repo.DB.Create(&reservation)
	if err.Error != nil {
		return reservations.Domain{}, err.Error
	}
	return reservation.ToDomain(), nil
}

func (Repo *reservationRepo) GetAllReservations(ctx context.Context) ([]reservations.Domain, error){
	var reservation []Reservation
	err := Repo.DB.Find(&reservation)
	if err.Error != nil {
		return []reservations.Domain{}, err.Error
	}
	return GetAllReservation(reservation), nil
}