package publishers_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iskandardevan/book-library/business/publishers"
	"github.com/iskandardevan/book-library/business/publishers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var publisherRepository = mocks.Repository{Mock:mock.Mock{}}
var publisherService publishers.PublishersUsecaseInterface
var publisherDomain publishers.Domain

func setup() {
	publisherService = publishers.NewUseCase(&publisherRepository, time.Hour*10)
	publisherDomain = publishers.Domain{
		Id: 	1,	
		Email:   "siteslagixixi@gmail.com",
		Name:    "sitesaja",
		Phone:   "9875465",
		Address: "awdawdawda adaw ",
	}
}

func TestAdd(t *testing.T) {
	setup()
	publisherRepository.On("AddPublisher", mock.Anything, mock.Anything).Return(publisherDomain, nil)
	publisherRepository.On(mock.Anything, mock.Anything).Return(publishers.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		publisher, err := publisherService.AddPublisher(context.Background(), publishers.Domain{
			Id:      1,
			Email:   "siteslagixixi@gmail.com",
			Name:    "sitesaja",
			Phone:   "9875465",
			Address: "awdawdawdaadaw",
		})

		assert.NoError(t, err)
		assert.Equal(t, publisherDomain, publisher)
	})
	t.Run("Test Case 2 | No Email", func(t *testing.T) {
		
		publisherRepository.On("AddPublisher", mock.Anything, mock.Anything).Return(publisherDomain, errors.New("email harus di isi")).Once() 
		publisher, err := publisherService.AddPublisher(context.Background(), publishers.Domain{
			Id 			:1,
			Email		:"",
			Name 		:"sitesaja",
			Phone		:"9875465",
			Address		:"digidaw",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, publisher)
	})
	t.Run("Test Case 3 | No Name", func(t *testing.T) {
		
		publisherRepository.On("AddPublisher", mock.Anything, mock.Anything).Return(publisherDomain, errors.New("nama harus di isi")).Once() 
		publisher, err := publisherService.AddPublisher(context.Background(), publishers.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"",
			Phone		:"9875465",
			Address		:"",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, publisher)
	})
	t.Run("Test Case 4 | No Phone", func(t *testing.T) {
		
		publisherRepository.On("AddPublisher", mock.Anything, mock.Anything).Return(publisherDomain, errors.New("nomer harus di isi")).Once() 
		publisher, err := publisherService.AddPublisher(context.Background(), publishers.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"sitesaja",
			Phone		:"",
			Address		:"digidaw",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, publisher)
	})
	t.Run("Test Case 5 | No Address", func(t *testing.T) {
		
		publisherRepository.On("AddPublisher", mock.Anything, mock.Anything).Return(publisherDomain, errors.New("alamat harus di isi")).Once() 
		publisher, err := publisherService.AddPublisher(context.Background(), publishers.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Name 		:"sitesaja",
			Phone		:"9875465",
			Address		:"",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, publisher)
	})
}