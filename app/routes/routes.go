package routes

import (
	_ "github.com/iskandardevan/book-library/app/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/iskandardevan/book-library/controllers/authors"
	"github.com/iskandardevan/book-library/controllers/books"
	"github.com/iskandardevan/book-library/controllers/publishers"
	"github.com/iskandardevan/book-library/controllers/reservations"
	"github.com/iskandardevan/book-library/controllers/users"
	"github.com/labstack/echo/v4"
)


type RouteControllerList struct {
	// JWTMiddleware middlewares.ConfigJWT
	UserController 		users.UserController
	AuthorController	authors.AuthorController
	PublisherController	publishers.PublisherController
	BookController		books.BookController
	ReservationController	reservations.ReservationController

}


func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	e.POST("user/register", ctrl.UserController.RegisterUser)
	e.POST("user/login", ctrl.UserController.LoginUser)
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.GET("users", ctrl.UserController.GetAllUsers)
	e.PUT("user/:id", ctrl.UserController.UpdateUserByID)
	e.DELETE("user/:id", ctrl.UserController.DeleteUserByID)


	e.POST("author/add", ctrl.AuthorController.AddAuthor)
	e.GET("authors", ctrl.AuthorController.GetAllAuthors)
	e.DELETE("author/:id", ctrl.AuthorController.Delete)

	e.POST("publisher/add", ctrl.PublisherController.AddPublisher)
	e.GET("publishers", ctrl.PublisherController.AddPublisher)
	e.DELETE("publisher/:id", ctrl.PublisherController.Delete)

	e.POST("book/add", ctrl.BookController.AddBook)
	e.GET("books", ctrl.BookController.GetAllBooks)
	e.DELETE("book/:id", ctrl.BookController.Delete)

	e.POST("reservation/add", ctrl.ReservationController.AddReservation)
	e.GET("reservations", ctrl.ReservationController.GetAllReservations)
	e.DELETE("reservation/:id", ctrl.ReservationController.Delete)
	
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
