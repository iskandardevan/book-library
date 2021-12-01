package reservations_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iskandardevan/book-library/business/reservations"
	"github.com/iskandardevan/book-library/business/reservations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reservationRepository = mocks.Repository{Mock: mock.Mock{}}
var reservationService reservations.ReservationUsecaseInterface
var reservationDomain reservations.Domain

var allreservationDomain []reservations.Domain

func setup() {
	reservationService = reservations.NewUseCase(&reservationRepository, time.Hour*10)
	reservationDomain = reservations.Domain{
		Id:      1,
		Book_ID: 2,
		User_ID: 3,
	}
}

func TestAdd(t *testing.T) {
	setup()
	reservationRepository.On("AddReservation", mock.Anything, mock.Anything).Return(reservationDomain, nil)
	reservationRepository.On(mock.Anything, mock.Anything).Return(reservations.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		reservation, err := reservationService.AddReservation(context.Background(), reservations.Domain{
			Id:      1,
			Book_ID: 2,
			User_ID: 3,
		})

		assert.NoError(t, err)
		assert.Equal(t, reservationDomain, reservation)
	})
	t.Run("Test Case 2 | No Book", func(t *testing.T) {
		
		reservationRepository.On("AddReservation", mock.Anything, mock.Anything).Return(reservationDomain, errors.New("buku harus di isi")).Once() 
		reservation, err := reservationService.AddReservation(context.Background(), reservations.Domain{
			Id:      1,
			Book_ID: 0,
			User_ID: 3,
			
		})
		assert.Error(t, err)
		assert.NotNil(t, reservation)
	})
	t.Run("Test Case 3 | No User", func(t *testing.T) {
		
		reservationRepository.On("AddReservation", mock.Anything, mock.Anything).Return(reservationDomain, errors.New("nama harus di isi")).Once() 
		reservation, err := reservationService.AddReservation(context.Background(), reservations.Domain{
			Id:      1,
			Book_ID: 2,
			User_ID: 0,
			
		})
		assert.Error(t, err)
		assert.NotNil(t, reservation)
	})
	t.Run("Test Case 4 | MT", func(t *testing.T) {
		
		reservationRepository.On("AddReservation", mock.Anything, mock.Anything).Return(reservationDomain, errors.New("tidak reservasi ada")).Once() 
		reservation, err := reservationService.AddReservation(context.Background(), reservations.Domain{
			Id:      0,
			Book_ID: 0,
			User_ID: 0,
		
		})
		
		assert.Error(t, err)
		assert.NotNil(t, reservation, reservationDomain)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success Search Reservation", func(t *testing.T) {
        setup()
        reservationRepository.On("GetAllReservations", mock.Anything).Return(allreservationDomain, nil).Once()
        data, err := reservationService.GetAllReservations(context.Background())

        assert.NoError(t, err)
        assert.Nil(t, data)
        assert.Equal(t, len(data), len(allreservationDomain))
    })

    t.Run("Test case 2 | Error Search Reservation(search empty)", func(t *testing.T) {
        setup()
        reservationRepository.On("GetAllReservations", mock.Anything, mock.Anything).Return([]reservations.Domain{}, errors.New("Reservation Not Found")).Once()
        data, err := reservationService.GetAllReservations(context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []reservations.Domain{})
    })
}