package publishers

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Email     	string
	Name		string
	Age			int
	Phone		string
	Address     string
}

type PublishersUsecaseInterface interface {
	
}

type PublishersRepoInterface interface {
	
}