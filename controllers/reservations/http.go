package reservations

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/reservations"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/reservations/request"
	"github.com/iskandardevan/book-library/controllers/reservations/response"
	"github.com/labstack/echo/v4"
)
type ReservationController struct {
	reservationUseCase reservations.ReservationUsecaseInterface
}

func NewReservationController(ReservationUseCase reservations.ReservationUsecaseInterface) *ReservationController{
	return &ReservationController{
		reservationUseCase: ReservationUseCase,
	}
}

func (reservationController *ReservationController) AddReservation (c echo.Context) error {
	req := request.AddReservationRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data reservations.Domain
	data, err = reservationController.reservationUseCase.AddReservation(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainReservation(data))

}

func (reservationController *ReservationController) GetAllReservations (c echo.Context) error {
	req := c.Request().Context()
	reservation, err := reservationController.reservationUseCase.GetAllReservations(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllReservation(reservation))

}
