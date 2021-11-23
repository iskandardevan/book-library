package repositoryuser

import (
	"errors"

	"github.com/iskandardevan/book-library/repository/users"
	"gorm.io/gorm"
)

type RepositoryUser interface {
	Create(data users.User) (users.User, error)
	IsEmailExist(email string) error
	Login(email string) (users.User, error)
	Update(data users.User) (users.User, error)
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryUser {
	return &repository{db: db}
}

func (r *repository) Create(data users.User) (users.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return users.User{}, err
	}

	return data, nil
}

func (r *repository) IsEmailExist(email string) error {
	user := new(users.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	return errors.New("email already exists")
}

func (r *repository) Login(email string) (users.User, error) {
	user := new(users.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return users.User{}, err
	}

	return *user, nil
}

func (r *repository) Update(data users.User) (users.User, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return users.User{}, err
	}

	return data, nil
}

func (r *repository) DeleteByID(id uint) error {
	user := new(users.User)
	user.ID = id
	return r.db.First(&user).Where("id = ?", user.ID).Delete(&user).Error
}