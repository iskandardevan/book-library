package authors

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/authors"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/authors/request"
	"github.com/iskandardevan/book-library/controllers/authors/response"
	"github.com/iskandardevan/book-library/helpers"
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

// AddAuthor author
// @Tags authors
// @Summary AddAuthor author
// @Description AddAuthor author
// @Accept  json
// @Produce  json
// @Param data body request.AddAuthorRequest true "data"
// @Success 200 {object} controllers.BaseResponse{data=response.AuthorResponse} "Add"
// @Router /author/add [POST]
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainAuthor(data))

}

func (authorController *AuthorController) GetAllAuthors(c echo.Context) error {
	req := c.Request().Context()
	author, err := authorController.authorUseCase.GetAllAuthors(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllAuthors(author))

}

func (authorController *AuthorController) Delete(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = authorController.authorUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}