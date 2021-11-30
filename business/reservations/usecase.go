package reservations

import (
	"context"
	"errors"
	"time"
)

type ReservationUseCase struct {
	repo ReservationRepoInterface
	ctx  time.Duration
}

func NewUseCase(reservationRepo ReservationRepoInterface, ctx  time.Duration) *ReservationUseCase{
	return &ReservationUseCase{
		repo:	reservationRepo,
		ctx: 	ctx,
	}
}

func (usecase *ReservationUseCase) AddReservation(ctx context.Context, domain Domain) (Domain, error){
	if domain.Book_ID == 0 {
		return Domain{}, errors.New("buku harus di isi")
	}
	if domain.User_ID == 0{
		return Domain{}, errors.New("nama harus di isi")
	}
	reservation, err := usecase.repo.AddReservation(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	return reservation, nil
}