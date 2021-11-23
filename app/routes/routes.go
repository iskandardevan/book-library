package routes

import (
	"github.com/labstack/echo"
)


func RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.GET("/users",)
	users.POST("", ctrlList.UserController.Register)
	m.
	users.POST("/login", ctrlList.UserController.Login)
}