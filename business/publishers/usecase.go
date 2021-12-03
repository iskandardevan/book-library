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
	if domain.Phone == "" {
		return Domain{}, errors.New("nomer hp harus di isi")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("alamat harus di isi")
	}

	publisher, err := usecase.repo.AddPublisher(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return publisher, nil
}

func (usecase *PublisherUseCase) GetAllPublishers(ctx context.Context) ([]Domain, error) {
	publisher, err := usecase.repo.GetAllPublishers(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada publisher")
	}
	return publisher, nil
}

func (usecase *PublisherUseCase) Delete(id uint, ctx context.Context) error{
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada publisher dengan ID tersebut")
	}
	return nil
}