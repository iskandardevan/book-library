package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint		`gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Name 	string
	Age			int
	Phone		string
	Password  string
	Token     string
}

type UserUsecaseInterface interface {
	LoginUser(email string, password string, ctx context.Context) (Domain, string, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	RegisterUser(ctx context.Context, domain Domain) (Domain, error)
	
}

type UserRepoInterface interface {
	GetByID(id uint) (Domain, error)
	RegisterUser(ctx context.Context, domain *Domain) (Domain, error)
	GetEmail(ctx context.Context, email string) (Domain, error)
}

