package authors

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

type AuthorUsecaseInterface interface {
	AddAuthor(ctx context.Context, domain Domain) (Domain, error)
	GetAllAuthors(ctx context.Context) ([]Domain, error)
	Delete(id uint, ctx context.Context)error
}

type AuthorRepoInterface interface {
	AddAuthor(ctx context.Context, domain Domain) (Domain, error)
	GetAllAuthors(ctx context.Context) ([]Domain, error)
	Delete(id uint, ctx context.Context)error
}

