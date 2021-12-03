package reservations

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/reservations"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/reservations/request"
	"github.com/iskandardevan/book-library/controllers/reservations/response"
	"github.com/iskandardevan/book-library/helpers"
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

// AddReservation reservation
// @Tags reservations
// @Summary AddReservation reservation
// @Description AddReservation reservation
// @Accept  json
// @Produce  json
// @Param data body request.AddReservationRequest true "data"
// @Success 200 {object} controllers.BaseResponse{data=response.ReservationResponse} "Add"
// @Router /reservation/add [POST]
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainReservation(data))

}

func (reservationController *ReservationController) GetAllReservations (c echo.Context) error {
	req := c.Request().Context()
	reservation, err := reservationController.reservationUseCase.GetAllReservations(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllReservation(reservation))

}

func (reservationController *ReservationController) Delete(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = reservationController.reservationUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}