package routes

import (
	"github.com/labstack/echo"
)


func RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.GET("/users",)
	m.
	users.POST("", ctrlList.UserController.Register)
	users.POST("/login", ctrlList.UserController.Login)
	
}