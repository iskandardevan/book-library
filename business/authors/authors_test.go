package authors_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iskandardevan/book-library/business/authors"
	"github.com/iskandardevan/book-library/business/authors/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var authorRepository = mocks.Repository{Mock:mock.Mock{}}
var authorService authors.AuthorUsecaseInterface
var authorDomain authors.Domain

func setup(){
	authorService = authors.NewUseCase(&authorRepository, time.Hour*10)
	authorDomain = authors.Domain{
		Email		:"siteslagixixi@gmail.com",
		Name 		:"sitesaja",
		Age			:231,
		Phone		:"9875465",
		Address		:"awdawdawda adaw ",
	}
}

func TestAdd(t *testing.T) {
	setup()
	authorRepository.On("AddAuthor", mock.Anything, mock.Anything).Return(authorDomain, nil)
	authorRepository.On(mock.Anything, mock.Anything).Return(authors.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		author, err := authorService.AddAuthor(context.Background(), authors.Domain{
		Id 			:1,
		Email		:"siteslagixixi@gmail.com",
		Name 		:"sitesaja",
		Age			:231,
		Phone		:"9875465",
		Address		:"awdawdawdaadaw",
		})

		assert.NoError(t, err)
		assert.Equal(t, authorDomain, author)
	})
	t.Run("Test Case 2 | No Email", func(t *testing.T) {
		
		authorRepository.On("AddAuthor", mock.Anything, mock.Anything).Return(authorDomain, errors.New("email harus di isi")).Once() 
		author, err := authorService.AddAuthor(context.Background(), authors.Domain{
			Id 			:1,
			Email		:"",
			Name 		:"sitesaja",
			Age			:231,
			Phone		:"9875465",
			Address		:"awdawdawdaadaw",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, author)
	})
	t.Run("Test Case 3 | No Name", func(t *testing.T) {
		
		authorRepository.On("AddAuthor", mock.Anything, mock.Anything).Return(authorDomain, errors.New("nama harus di isi")).Once() 
		author, err := authorService.AddAuthor(context.Background(), authors.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"",
			Age			:231,
			Phone		:"9875465",
			Address		:"awdawdawdaadaw",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, author)
	})
	t.Run("Test Case 4 | No Age", func(t *testing.T) {
		
		authorRepository.On("AddAuthor", mock.Anything, mock.Anything).Return(authorDomain, errors.New("umur harus di isi")).Once() 
		author, err := authorService.AddAuthor(context.Background(), authors.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"sitesaja",
			Age			:0,
			Phone		:"9875465",
			Address		:"awdawdawdaadaw",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, author)
	})
	t.Run("Test Case 5 | No Phone", func(t *testing.T) {
		
		authorRepository.On("AddAuthor", mock.Anything, mock.Anything).Return(authorDomain, errors.New("nomer hp harus di isi")).Once() 
		author, err := authorService.AddAuthor(context.Background(), authors.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"sitesaja",
			Age			:231,
			Phone		:"",
			Address		:"awdawdawdaadaw",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, author)
	})
	t.Run("Test Case 6 | No Address", func(t *testing.T) {
		
		authorRepository.On("AddAuthor", mock.Anything, mock.Anything).Return(authorDomain, errors.New("alamat harus di isi")).Once() 
		author, err := authorService.AddAuthor(context.Background(), authors.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"sitesaja",
			Age			:231,
			Phone		:"9875465",
			Address		:"",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, author)
	})
}