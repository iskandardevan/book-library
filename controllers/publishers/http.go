package publishers

import (
	"net/http"

	"github.com/iskandardevan/book-library/business/publishers"
	"github.com/iskandardevan/book-library/controllers"
	"github.com/iskandardevan/book-library/controllers/publishers/request"
	"github.com/iskandardevan/book-library/controllers/publishers/response"
	"github.com/iskandardevan/book-library/helpers"
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

// AddPublisher publisher
// @Tags publishers
// @Summary AddPublisher publisher
// @Description AddPublisher publisher
// @Accept  json
// @Produce  json
// @Param data body request.AddPublisherRequest true "data"
// @Success 200 {object} controllers.BaseResponse{data=response.PublisherResponse} "Add"
// @Router /publisher/add [POST]
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
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainPublisher(data))

}

func (publisherController *PublisherController) GetAllPublishers(c echo.Context) error {
	req := c.Request().Context()
	publisher, err := publisherController.publisherUseCase.GetAllPublishers(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllPublishers(publisher))

}

func (publisherController *PublisherController) Delete(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = publisherController.publisherUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}