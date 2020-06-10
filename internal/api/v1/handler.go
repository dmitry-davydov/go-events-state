package v1

import (
	"net/http"

	"github.com/dmitry-davydov/go-events-state/internal/api/v1/request"
	"github.com/dmitry-davydov/go-events-state/internal/api/v1/response"
	"github.com/dmitry-davydov/go-events-state/internal/models"
	"github.com/dmitry-davydov/go-events-state/internal/storage"

	echo "github.com/labstack/echo/v4"
)

func AttachRoutes(group *echo.Group, storage *storage.Storage) {

	group.GET("/events", func(c echo.Context) error {
		all := storage.All()
		var allResponse []response.Event

		for _, e := range all {
			var resp response.Event
			resp.Map(e)
			allResponse = append(allResponse, resp)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"data": allResponse,
		})
	})

	group.GET("/event/:eventId", func(c echo.Context) error {
		event := storage.Event(c.Param("eventId"))
		if event == nil {
			return c.NoContent(http.StatusNotFound)
		}

		var resp response.Event
		resp.Map(event)

		return c.JSON(http.StatusOK, echo.Map{
			"data": resp,
		})
	})

	group.POST("/event", func(c echo.Context) error {
		ev := models.NewEvent()

		storage.SetEvent(ev)

		var responseDto response.Event
		responseDto.Map(ev)

		return c.JSON(http.StatusOK, echo.Map{
			"data": responseDto,
		})

	})

	group.PUT("/event/:eventId", func(c echo.Context) error {

		requestObj := new(request.Event)
		if err := c.Bind(requestObj); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if !requestObj.IsValid() {
			return c.NoContent(http.StatusBadRequest)
		}

		event := storage.Event(c.Param("eventId"))
		if event == nil {
			return c.NoContent(http.StatusNotFound)
		}

		event.SetState(requestObj.State)
		storage.SetEvent(event)

		var responseObj response.Event
		responseObj.Map(event)
		return c.JSON(http.StatusOK, echo.Map{
			"data": responseObj,
		})
	})

	group.DELETE("/event/:eventId", func(c echo.Context) error {
		if !storage.Delete(c.Param("eventId")) {
			return c.NoContent(http.StatusNotFound)
		}

		return c.NoContent(http.StatusOK)
	})

}
