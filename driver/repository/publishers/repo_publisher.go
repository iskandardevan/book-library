package publishers

import (
	"context"

	"github.com/iskandardevan/book-library/business/publishers"
	"gorm.io/gorm"
)

type publisherRepo struct {
	db *gorm.DB
}

func NewPublisherRepo(db *gorm.DB) *publisherRepo {
	return &publisherRepo{db: db}
}

func (Repo *publisherRepo) RegisterPublisher(ctx context.Context, domain *publishers.Domain) (publishers.Domain, error) {
	publisher := Publisher{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Phone		:domain.Phone,
		Address		:domain.Address,
	}
	err := Repo.db.Create(&publisher)
	if err.Error != nil {
		return publishers.Domain{}, err.Error
	}
	return publisher.ToDomain(), nil
} 