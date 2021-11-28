package authors

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/authors"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/authors/request"
	"github.com/iskandardevan/book-library/controllers/authors/response"
	"github.com/labstack/echo/v4"
)

type AuthorController struct {
	authorUseCase authors.AuthorUsecaseInterface
}

func NewAuthorController(AuthorUseCase authors.AuthorUsecaseInterface) *AuthorController{
	return &AuthorController{
		authorUseCase: AuthorUseCase,
	}
}

func (authorController *AuthorController) AddAuthor (c echo.Context) error {
	req := request.AddAuthorRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data authors.Domain
	data, err = authorController.authorUseCase.AddAuthor(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainAuthor(data))

}

