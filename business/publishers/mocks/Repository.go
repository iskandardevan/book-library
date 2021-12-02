package mocks

import (
	"context"

	"github.com/iskandardevan/book-library/business/publishers"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) AddPublisher(ctx context.Context, domain publishers.Domain) (publishers.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 publishers.Domain
	if rf, ok := ret.Get(0).(func(context.Context, publishers.Domain) publishers.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(publishers.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, publishers.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}