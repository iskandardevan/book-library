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
	req := request.RegisterUserRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data users.Domain
	data, err = userController.userUseCase.RegisterUser(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))

}


func (userController *UserController) LoginUser (c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	req := request.UserLoginRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	login, token, err = userController.userUseCase.LoginUser(req.Email, req.Password, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.UserLogin(login, token))
}

func (userController *UserController) GetByID (c echo.Context) error{
	var getid users.Domain
	var err error
	ctx := c.Request().Context()
	req := request.GetByID{}
	err = c.Bind(&req)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	getid, err = userController.userUseCase.GetByID(req.Id, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.GetByID(getid))

}
