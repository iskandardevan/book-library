package publishers

import (
	"context"
	"errors"

	"github.com/iskandardevan/book-library/business/publishers"
	"gorm.io/gorm"
)

type publisherRepo struct {
	DB *gorm.DB
}

func NewPublisherRepo(DB *gorm.DB) *publisherRepo {
	return &publisherRepo{DB: DB}
}

func (Repo *publisherRepo) AddPublisher(ctx context.Context, domain publishers.Domain) (publishers.Domain, error) {
	publisher := Publisher{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Phone		:domain.Phone,
		Address		:domain.Address,
	}
	err := Repo.DB.Create(&publisher)
	if err.Error != nil {
		return publishers.Domain{}, err.Error
	}
	return publisher.ToDomain(), nil
} 

func (Repo *publisherRepo) GetAllPublishers(ctx context.Context) ([]publishers.Domain, error){
	var publisher []Publisher
	err := Repo.DB.Find(&publisher)
	if err.Error != nil {
		return []publishers.Domain{}, err.Error
	}
	return GetAllPublishers(publisher), nil
}

func (Repo *publisherRepo) Delete(id uint, ctx context.Context) error{
	publisher := Publisher{}
	err := Repo.DB.Delete(&publisher, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}