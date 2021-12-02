package routes

import (
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
	e.PUT("user/:id", ctrl.UserController.UpdateUserByID)
	e.DELETE("user/:id", ctrl.UserController.DeleteUserByID)


	e.POST("author/add", ctrl.AuthorController.AddAuthor)

	e.POST("publisher/add", ctrl.PublisherController.AddPublisher)

	e.POST("book/add", ctrl.BookController.AddBook)
	e.GET("books", ctrl.BookController.GetAllBooks)

	e.POST("reservation/add", ctrl.ReservationController.AddReservation)
	e.GET("reservations", ctrl.ReservationController.GetAllReservations)
	
	
}
