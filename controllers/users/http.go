package users

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/users"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/users/request"
	"github.com/iskandardevan/book-library/controllers/users/response"
	"github.com/labstack/echo/v4"
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
	req := request.RegisterUser{}
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

func (userController *UserController) GetAllUsers (c echo.Context) error {
	req := c.Request().Context()
	data, err := userController.userUseCase.GetAllUsers(req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllUsers(data))

}

// func (userController *UserController) UpdateUser (c echo.Context) error {
	
// }

// func (userController *UserController) DeleteUser (c echo.Context) error {
	
// }