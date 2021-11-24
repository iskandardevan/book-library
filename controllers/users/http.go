package users

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/users"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/users/request"
	"github.com/iskandardevan/book-library/controllers/users/response"
	"github.com/labstack/echo"
)


type UserController struct {
	userUseCase users.UserUsecaseInterface
}

func NewUserController(UserUseCase users.UserUsecaseInterface) *UserController{
	return &UserController{
		userUseCase: UserUseCase,
	}
}

func (userController *UserController) RegisterUser (c echo.Context) error {
	req := request.CreateUser{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	var data users.Domain
	data, err = userController.userUseCase.RegisterUser(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))

}
