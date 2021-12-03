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

var allpublisherDomain []publishers.Domain

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

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success Search publisher", func(t *testing.T) {
        setup()
        publisherRepository.On("GetAllPublishers", mock.Anything).Return(allpublisherDomain, nil).Once()
        data, err := publisherService.GetAllPublishers(context.Background())

        assert.NoError(t, err)
        assert.Nil(t, data)
        assert.Equal(t, len(data), len(allpublisherDomain))
    })

    t.Run("Test case 2 | Error Search publisher(search empty)", func(t *testing.T) {
        setup()
        publisherRepository.On("GetAllPublishers", mock.Anything, mock.Anything).Return([]publishers.Domain{}, errors.New("tidak ada publisher")).Once()
        data, err := publisherService.GetAllPublishers(context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []publishers.Domain{})
    })
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		publisherRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := publisherService.Delete(publisherDomain.Id, context.Background() )

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		publisherRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("tidak ada publisher dengan ID tersebut")).Once()
		err := publisherService.Delete(publisherDomain.Id, context.Background())

		assert.Equal(t, err, errors.New("tidak ada publisher dengan ID tersebut"))
		assert.Error(t, err)
	})
}