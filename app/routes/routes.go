package routes

import (
	"github.com/iskandardevan/book-library/controllers/authors"
	"github.com/iskandardevan/book-library/controllers/books"
	"github.com/iskandardevan/book-library/controllers/publishers"
	"github.com/iskandardevan/book-library/controllers/users"
	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	// JWTMiddleware middlewares.ConfigJWT
	UserController 		users.UserController
	AuthorController	authors.AuthorController
	PublisherController	publishers.PublisherController
	BookController		books.BookController
}


func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	e.POST("user/register", ctrl.UserController.RegisterUser)
	e.POST("user/login", ctrl.UserController.LoginUser)
	
	e.POST("author/add", ctrl.AuthorController.AddAuthor)
	e.POST("publisher/add", ctrl.PublisherController.AddPublisher)
	e.POST("book/add", ctrl.BookController.AddBook)
	// e.GET("users", ctrl.UserController.GetAllUsers)
	
	
}
