package reservations

import (
	"context"
	"time"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/iskandardevan/book-library/business/users"
	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Book     	books.Domain
	Book_ID		uint
	User		users.Domain
	User_ID		uint
	Start		time.Time
	End			time.Time
}

type ReservationUsecaseInterface interface {
	AddReservation(ctx context.Context, domain Domain) (Domain, error)
	
}

type ReservationRepoInterface interface {
	AddReservation(ctx context.Context, domain *Domain) (Domain, error)
	
}
