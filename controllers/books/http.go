package books

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/books/request"
	"github.com/iskandardevan/book-library/controllers/books/response"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookUseCase books.BooksUsecaseInterface
}

func NewBookController(BookUseCase books.BooksUsecaseInterface) *BookController{
	return &BookController{
		bookUseCase: BookUseCase,
	}
}

func (bookController *BookController) AddBook (c echo.Context) error {
	req := request.AddBookRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data books.Domain
	data, err = bookController.bookUseCase.AddBook(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBook(data))

}

func (bookController *BookController) GetAllBooks (c echo.Context) error {
	req := c.Request().Context()
	book, err := bookController.bookUseCase.GetAllBooks(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllBook(book))

}
