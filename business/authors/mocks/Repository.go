package mocks

import (
	"context"

	"github.com/iskandardevan/book-library/business/authors"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) AddAuthor(ctx context.Context, domain authors.Domain) (authors.Domain, error){
	ret := _m.Called(ctx, domain)

	var r0 authors.Domain
	if rf, ok := ret.Get(0).(func(context.Context, authors.Domain) authors.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(authors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, authors.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *Repository) GetAllAuthors(ctx context.Context) ([]authors.Domain, error) {
    ret := _m.Called(ctx)

    var r0 []authors.Domain
    if rf, ok := ret.Get(0).(func(context.Context) []authors.Domain); ok {
        r0 = rf(ctx)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).([]authors.Domain)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context) error); ok {
        r1 = rf(ctx)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *Repository) Delete (id uint, ctx context.Context ) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}