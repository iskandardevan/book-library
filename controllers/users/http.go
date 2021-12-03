package users

import (
	"net/http"
	"strconv"

	"github.com/iskandardevan/book-library/business/users"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/users/request"
	"github.com/iskandardevan/book-library/controllers/users/response"
	"github.com/iskandardevan/book-library/helpers"
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

// RegisterUser user
// @Tags users
// @Summary RegisterUser user
// @Description RegisterUser user
// @Accept  json
// @Produce  json
// @Param data body request.RegisterUserRequest true "data"
// @Success 200 {object} controllers.BaseResponse{data=response.UserResponse} "Register"
// @Router /user/register [POST]
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
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := userController.userUseCase.GetByID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))
}

func (userController *UserController) GetAllUsers (c echo.Context) error {
	req := c.Request().Context()
	user, err := userController.userUseCase.GetAllUsers(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllUsers(user))

}

func (userController *UserController) UpdateUserByID (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.RegisterUserRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := userController.userUseCase.UpdateUserByID(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))

}

func (userController *UserController) DeleteUserByID (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = userController.userUseCase.DeleteUserByID(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}