package reservations

import (
	"time"

	"github.com/iskandardevan/book-library/business/reservations"
	"github.com/iskandardevan/book-library/driver/repository/books"
	"github.com/iskandardevan/book-library/driver/repository/users"
	"gorm.io/gorm"
)

type Reservation struct {
	Id        uint 			`gorm:"primaryKey"`
	Book      books.Book	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Book_ID		uint	
	User      users.User 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User_ID		uint
	Start     time.Time
	End       time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (reservation *Reservation) ToDomain() reservations.Domain {
	return reservations.Domain{
		Id 			:reservation.Id,
		Book_ID     :reservation.Book_ID,
		User_ID     :reservation.User_ID,
		Start     	:reservation.Start,
		End       	:reservation.End,
		CreatedAt 	:reservation.CreatedAt,
		UpdatedAt 	:reservation.UpdatedAt,
	}
}

func FromDomain(domain reservations.Domain) Reservation  {
	return Reservation{
		Id 			:domain.Id,
		Book_ID     :domain.Book_ID,
		User_ID   	:domain.User_ID,
		Start     	:domain.Start,
		End       	:domain.End,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}

func GetAllReservation(data []Reservation) []reservations.Domain{
	res := []reservations.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}