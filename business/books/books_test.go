package books_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/iskandardevan/book-library/business/books/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bookRepository = mocks.Repository{Mock: mock.Mock{}}
var bookService books.BooksUsecaseInterface
var bookDomain books.Domain

var allbookDomain []books.Domain

func setup() {
	bookService = books.NewUseCase(&bookRepository, time.Hour*10)
	bookDomain = books.Domain{
		Id:      1,
			Title:            "Book.Title",
			Edition:          1,
			Author_ID:        2,
			Publisher_ID:     3,
			Publication_Year: "2021",
	}
}


func TestAdd(t *testing.T) {
	setup()
	bookRepository.On("AddBook", mock.Anything, mock.Anything).Return(bookDomain, nil)
	bookRepository.On(mock.Anything, mock.Anything).Return(books.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		book, err := bookService.AddBook(context.Background(), books.Domain{
			Id:      1,
			Title:            "Book.Title",
			Edition:          1,
			Author_ID:        2,
			Publisher_ID:     3,
			Publication_Year: "2021",
		})

		assert.NoError(t, err)
		assert.Equal(t, bookDomain, book)
	})

	t.Run("Test Case 2 | No Title", func(t *testing.T) {
		
		bookRepository.On("AddBook", mock.Anything, mock.Anything).Return(bookDomain, errors.New("title harus di isi")).Once() 
		book, err := bookService.AddBook(context.Background(), books.Domain{
			Id:      1,
			Title:            "",
			Edition:          1,
			Author_ID:        2,
			Publisher_ID:     3,
			Publication_Year: "2021",
		})
		assert.Error(t, err)
		assert.NotNil(t, book)
	})

	t.Run("Test Case 3 | No Edition", func(t *testing.T) {
		
		bookRepository.On("AddBook", mock.Anything, mock.Anything).Return(bookDomain, errors.New("edition harus di isi")).Once() 
		book, err := bookService.AddBook(context.Background(), books.Domain{
			Id:      1,
			Title:            "Book.Title",
			Edition:          0,
			Author_ID:        2,
			Publisher_ID:     3,
			Publication_Year: "2021",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, book)
	})

	t.Run("Test Case 4 | No Publication Year", func(t *testing.T) {
		
		bookRepository.On("AddBook", mock.Anything, mock.Anything).Return(bookDomain, errors.New("tahun terbit harus di isi")).Once() 
		basing, err := bookService.AddBook(context.Background(), books.Domain{
			Id:      1,
			Title:            "Book.Title",
			Edition:          1,
			Author_ID:        2,
			Publisher_ID:     3,
			Publication_Year: "",
			
		})
		assert.Error(t, err)
		assert.NotNil(t, basing)
	})
	
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success Search book", func(t *testing.T) {
        setup()
        bookRepository.On("GetAllBooks", mock.Anything).Return(allbookDomain, nil).Once()
        data, err := bookService.GetAllBooks(context.Background())

        assert.NoError(t, err)
        assert.Nil(t, data)
        assert.Equal(t, len(data), len(allbookDomain))
    })

    t.Run("Test case 2 | Error Search book(search empty)", func(t *testing.T) {
        setup()
        bookRepository.On("GetAllBooks", mock.Anything, mock.Anything).Return([]books.Domain{}, errors.New("book Not Found")).Once()
        data, err := bookService.GetAllBooks(context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []books.Domain{})
    })
}