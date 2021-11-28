package publishers

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint			`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt 	`gorm:"index"`
	Email     	string
	Name		string
	Age			int
	Phone		string
	Address     string
}

type PublishersUsecaseInterface interface {
	AddPublisher(ctx context.Context, domain Domain) (Domain, error)
}

type PublishersRepoInterface interface {
	AddPublisher(ctx context.Context, domain *Domain) (Domain, error)
}