package reservations

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Book     	string
	User		string
	Start		time.Time
	End			time.Time
}
