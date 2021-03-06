package users_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iskandardevan/book-library/app/middlewares"
	"github.com/iskandardevan/book-library/business/users"
	"github.com/iskandardevan/book-library/business/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = mocks.Repository{Mock:mock.Mock{}}
var userService users.UserUsecaseInterface
var token string 
var userDomain users.Domain

var alluserDomain []users.Domain

func setup(){
	userService = users.NewUseCase(&userRepository, time.Hour*10, &middlewares.ConfigJWT{})
	userDomain = users.Domain{
		Id 			:1,
		Email		:"sites@gmail.com",
		Name 		:"sites",
		Age			:21,
		Phone		:"198198198",
		Password	:"qqwerty",
	}
	token = "bvasingdaisnya"
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("RegisterUser", mock.Anything, mock.Anything).Return(userDomain, nil)
	userRepository.On("GetEmail", mock.Anything, mock.Anything).Return(users.Domain{}, nil)
	t.Run("Test Case 1 | Success Register", func(t *testing.T) {
		user, err := userService.RegisterUser(context.Background(), users.Domain{
			Id 			:1,
			Email		:"sites@gmail.com",
			Name 		:"sites",
			Age			:21,
			Phone		:"198198198",
			Password	:"qqwerty",
		})

		assert.NoError(t, err)
		assert.Equal(t, userDomain, user)
	})
	t.Run("Test Case 2 | No Email", func(t *testing.T) {
		
		userRepository.On("RegisterUser", mock.Anything, mock.Anything).Return(userDomain, errors.New("email belum di isi")).Once() 
		user, err := userService.RegisterUser(context.Background(), users.Domain{
			Id 			:1,
			Email		:"",
			Password	:"9875465",
		})
		assert.Error(t, err)
		assert.NotNil(t, user)
	})
	t.Run("Test Case 3 | No Password", func(t *testing.T) {
		
		userRepository.On("RegisterUser", mock.Anything, mock.Anything).Return(userDomain, errors.New("password belum di isi")).Once() 
		user, err := userService.RegisterUser(context.Background(), users.Domain{
			Id 			:1,
			Email		:"siteslagixixi@gmail.com",
			Password	:"",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, user)
	})
}



func TestUpdate(t *testing.T) {
	t.Run("Test case 1 | Success Update", func(t *testing.T) {
		setup()
		userRepository.On("UpdateUserByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.UpdateUserByID(userDomain.Id, context.Background(), userDomain )

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update", func(t *testing.T) {
		setup()
		userRepository.On("UpdateUserByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, errors.New("tidak ada user dengan ID tersebut")).Once()
		data, err := userService.UpdateUserByID(userDomain.Id, context.Background(), userDomain)

		assert.Equal(t, data, users.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		userRepository.On("DeleteUserByID", mock.Anything, mock.Anything).Return(nil).Once()
		err := userService.DeleteUserByID(userDomain.Id, context.Background() )

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		userRepository.On("DeleteUserByID", mock.Anything, mock.Anything).Return(errors.New("tidak ada user dengan ID tersebut")).Once()
		err := userService.DeleteUserByID(userDomain.Id, context.Background())

		assert.Equal(t, err, errors.New("tidak ada user dengan ID tersebut"))
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Success Login", func(t *testing.T) {
		setup()
		userRepository.On("GetEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		user, token, err := userService.LoginUser(userDomain.Email, "123", context.Background())

		assert.NotNil(t, token)
		assert.NoError(t, err)
		assert.Equal(t, user, users.Domain{})
	})

	t.Run("Test Case 2 | Cannot Login ", func(t *testing.T) {
		data, token, err := userService.LoginUser(userDomain.Email, "", context.Background())

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("Test Case 3 | Cannot Login ", func(t *testing.T) {
		setup()
		userRepository.On("GetEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("tidak ada user dengan ID tersebut")).Once()
		data, token, err := userService.LoginUser(userDomain.Email, "qqwerty", context.Background())

		assert.Equal(t, users.Domain{}, data, token, "")
		assert.NoError(t, err)
	})
	
	t.Run("Test Case 4 | No Email ", func(t *testing.T) {
		setup()
		userRepository.On("GetEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("email belum di isi")).Once()
		data, token, err := userService.LoginUser("", userDomain.Password, context.Background())

		assert.Equal(t, users.Domain{}, data, token, "")
		assert.Error(t, err)
	})

	
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetByID", func(t *testing.T) {
		setup()
		userRepository.On("GetByID", mock.AnythingOfType("uint") ,mock.Anything ).Return(userDomain, nil).Once()
		user, err := userService.GetByID(userDomain.Id, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("Test case 2 | Error GetByID(user Id = 0)", func(t *testing.T) {
		setup()
		userDomain.Id = 0
		userRepository.On("GetByID", mock.AnythingOfType("uint") ,mock.Anything ).Return(userDomain, nil).Once()
		data, err := userService.GetByID(userDomain.Id, context.Background())

		assert.Error(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, users.Domain{})
	})

	t.Run("Test case 3 | Error GetByID", func(t *testing.T) {
		setup()
		userRepository.On("GetByID", mock.AnythingOfType("uint") ,mock.Anything ).Return(users.Domain{}, nil).Once()
		data, err := userService.GetByID(3, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, users.Domain{})
	})

}

func TestGetAllUsers(t *testing.T) {
	t.Run("Test case 1 | Success Search user", func(t *testing.T) {
        setup()
        userRepository.On("GetAllUsers", mock.Anything).Return(alluserDomain, nil).Once()
        data, err := userService.GetAllUsers(context.Background())

        assert.NoError(t, err)
        assert.Nil(t, data)
        assert.Equal(t, len(data), len(alluserDomain))
    })

    t.Run("Test case 2 | Error Search user", func(t *testing.T) {
        setup()
        userRepository.On("GetAllUsers", mock.Anything, mock.Anything).Return([]users.Domain{}, errors.New("user Not Found")).Once()
        data, err := userService.GetAllUsers(context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []users.Domain{})
    })
}