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
	reservation, err := usecase.repo.AddReservation(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return reservation, nil
}

func (usecase *ReservationUseCase) GetAllReservations(ctx context.Context) ([]Domain, error) {
	reservation, err := usecase.repo.GetAllReservations(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada reservasi")
	}
	return reservation, nil
}

func (usecase *ReservationUseCase) Delete(id uint, ctx context.Context) error{
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada Reservation dengan ID tersebut")
	}
	return nil
}