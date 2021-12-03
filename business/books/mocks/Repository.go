package mocks

import (
	"context"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) AddBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 books.Domain
	if rf, ok := ret.Get(0).(func(context.Context, books.Domain) books.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(books.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, books.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *Repository) GetAllBooks(ctx context.Context) ([]books.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []books.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []books.Domain); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).([]books.Domain)
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