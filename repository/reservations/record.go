package reservations

import (
	"time"

	"github.com/iskandardevan/book-library/business/reservations"
	"gorm.io/gorm"
)

type Reservation struct {
	Id        uint `gorm:"primaryKey"`
	Book      string
	User      string `gorm:"unique"`
	Start     time.Time
	End       time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (reservation *Reservation) ToDomain() reservations.Domain {
	return reservations.Domain{
		Id 			:reservation.Id,
		Book     	:reservation.Book,
		User      	:reservation.User,
		Start     	:reservation.Start,
		End       	:reservation.End,
		CreatedAt 	:reservation.CreatedAt,
		UpdatedAt 	:reservation.UpdatedAt,
	}
}

func FromDomain(domain reservations.Domain) Reservation  {
	return Reservation{
		Id 			:domain.Id,
		Book     	:domain.Book,
		User      	:domain.User,
		Start     	:domain.Start,
		End       	:domain.End,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}