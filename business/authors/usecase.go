package authors

import (
	"context"
	"errors"
	"time"
)

type AuthorUseCase struct {
	repo AuthorRepoInterface
	ctx  time.Duration
}

func NewUseCase(authorRepo AuthorRepoInterface, ctx time.Duration) *AuthorUseCase {
	return &AuthorUseCase{
		repo: authorRepo,
		ctx:  ctx,
	}
}

func (usecase *AuthorUseCase) AddAuthor(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email harus di isi")
	}
	if domain.Name == "" {
		return Domain{}, errors.New("nama harus di isi")
	}
	if domain.Age == 0 {
		return Domain{}, errors.New("umur harus di isi")
	}
	if domain.Phone == "" {
		return Domain{}, errors.New("nomer hp harus di isi")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("alamat harus di isi")
	}

	author, err := usecase.repo.AddAuthor(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return author, nil
}

func (usecase *AuthorUseCase) GetAllAuthors(ctx context.Context) ([]Domain, error) {
	author, err := usecase.repo.GetAllAuthors(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada buku")
	}
	return author, nil
}

func (usecase *AuthorUseCase) Delete(id uint, ctx context.Context) error{
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada Author dengan ID tersebut")
	}
	return nil
}