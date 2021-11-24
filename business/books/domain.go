package books

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id               	uint
	CreatedAt        	time.Time
	UpdatedAt        	time.Time
	DeletedAt        	gorm.DeletedAt `gorm:"index"`
	Title            	string
	Edition          	int
	Author				string
	Publisher			string
	Publication_Year 	time.Time
}
