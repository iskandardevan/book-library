package mocks

import (
	"context"

	"github.com/iskandardevan/book-library/business/reservations"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) AddReservation(ctx context.Context, domain reservations.Domain) (reservations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 reservations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, reservations.Domain) reservations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(reservations.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, reservations.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *Repository) GetAllReservations(ctx context.Context) ([]reservations.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []reservations.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []reservations.Domain); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).([]reservations.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}