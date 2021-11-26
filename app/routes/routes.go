package routes

import (
	"github.com/iskandardevan/book-library/controllers/users"
	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	// JWTMiddleware middlewares.ConfigJWT
	UserController users.UserController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	e.POST("user/register", ctrl.UserController.RegisterUser)
	e.POST("user/login", ctrl.UserController.LoginUser)
	// e.GET("users", ctrl.UserController.GetAllUsers)
	// users := e.Group("users")
	// users.GET("/users",)
	// users.POST("", ctrlList.UserController.Register)
	// users.POST("/login", ctrlList.UserController.Login)
	
}

