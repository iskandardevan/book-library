package books

import (
	"context"
	"errors"
	"time"
)

type BookUseCase struct {
	repo BooksRepoInterface
	ctx  time.Duration
}

func NewUseCase(bookRepo BooksRepoInterface, ctx time.Duration) *BookUseCase {
	return &BookUseCase{
		repo: bookRepo,
		ctx:  ctx,
	}
}

func (usecase *BookUseCase) AddBook(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, errors.New("title harus di isi")
	}
	if domain.Edition == 0 {
		return Domain{}, errors.New("edition harus di isi")
	}
	if domain.Publication_Year == "" {
		return Domain{}, errors.New("tahun terbit harus di isi")
	}
	book, err := usecase.repo.RegisterBook(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return book, nil
}

func (usecase *BookUseCase) GetAllBooks(ctx context.Context) ([]Domain, error) {
	book, err := usecase.repo.GetAllBooks(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada buku")
	}
	return book, nil
}