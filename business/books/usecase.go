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
	// if domain.Publication_Year == 0 {
	// 	return Domain{}, errors.New("umur harus di isi")
	// }
	book, err := usecase.repo.RegisterBook(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return book, nil
}