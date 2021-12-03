package authors

import (
	"context"
	"errors"

	"github.com/iskandardevan/book-library/business/authors"
	"gorm.io/gorm"
)

type authorRepo struct {
	DB *gorm.DB
}

func NewAuthorRepo(DB *gorm.DB) authors.AuthorRepoInterface {
	return &authorRepo{DB: DB}
}

func (Repo *authorRepo) AddAuthor(ctx context.Context, domain authors.Domain) (authors.Domain, error) {
	author := Author{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Age			:domain.Age,
		Phone		:domain.Phone,
		Address		:domain.Address,
	}
	err := Repo.DB.Create(&author)
	if err.Error != nil {
		return authors.Domain{}, err.Error
	}
	return author.ToDomain(), nil
}

func (Repo *authorRepo) GetAllAuthors(ctx context.Context) ([]authors.Domain, error){
	var author []Author
	err := Repo.DB.Find(&author)
	if err.Error != nil {
		return []authors.Domain{}, err.Error
	}
	return GetAllAuthors(author), nil
}

func (Repo *authorRepo) Delete(id uint, ctx context.Context) error{
	author := Author{}
	err := Repo.DB.Delete(&author, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}