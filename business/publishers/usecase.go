package publishers

import (
	"context"
	"errors"
	"time"
)

type PublisherUseCase struct {
	repo PublishersRepoInterface
	ctx  time.Duration
}

func NewUseCase(publisherRepo PublishersRepoInterface, ctx time.Duration) *PublisherUseCase {
	return &PublisherUseCase{
		repo: publisherRepo,
		ctx:  ctx,
	}
}

func (usecase *PublisherUseCase) AddPublisher(ctx context.Context, domain Domain) (Domain, error) {
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

	publisher, err := usecase.repo.RegisterPublisher(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return publisher, nil
}
