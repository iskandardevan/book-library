package reservations

import (
	"context"

	"github.com/iskandardevan/book-library/business/reservations"
	"gorm.io/gorm"
)

type reservationRepo struct {
	db *gorm.DB
}

func NewReservationRepo(db *gorm.DB) *reservationRepo {
	return &reservationRepo{db: db}
}

func (Repo *reservationRepo) Registerreservation(ctx context.Context, domain *reservations.Domain) (reservations.Domain, error) {
	reservation := Reservation{
		Id:    domain.Id,
		Book:  domain.Book,
		User:  domain.User,
		Start: domain.Start,
		End:   domain.End,
	}
	err := Repo.db.Create(&reservation)
	if err.Error != nil {
		return reservations.Domain{}, err.Error
	}
	return reservation.ToDomain(), nil
}