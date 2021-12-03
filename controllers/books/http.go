package books

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/books"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/books/request"
	"github.com/iskandardevan/book-library/controllers/books/response"
	"github.com/iskandardevan/book-library/helpers"
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

// AddBook book
// @Tags books
// @Summary AddBook book
// @Description AddBook book
// @Accept  json
// @Produce  json
// @Param data body request.AddBookRequest true "data"
// @Success 200 {object} controllers.BaseResponse{data=response.BookResponse} "Add"
// @Router /book/add [POST]
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBook(data))

}

func (bookController *BookController) GetAllBooks (c echo.Context) error {
	req := c.Request().Context()
	book, err := bookController.bookUseCase.GetAllBooks(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllBook(book))

}

func (bookController *BookController) Delete(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = bookController.bookUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}