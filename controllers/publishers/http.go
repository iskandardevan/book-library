package publishers

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/publishers"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/publishers/request"
	"github.com/iskandardevan/book-library/controllers/publishers/response"
	"github.com/labstack/echo/v4"
)

type PublisherController struct {
	publisherUseCase publishers.PublishersUsecaseInterface
}

func NewPublisherController(PublisherUseCase publishers.PublishersUsecaseInterface) *PublisherController {
	return &PublisherController{
		publisherUseCase: PublisherUseCase,
	}
}

func (publisherController *PublisherController) AddPublisher(c echo.Context) error {
	req := request.AddPublisherRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data publishers.Domain
	data, err = publisherController.publisherUseCase.AddPublisher(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainPublisher(data))

}