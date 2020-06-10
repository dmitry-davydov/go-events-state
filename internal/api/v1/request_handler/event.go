package request_handler

import (


	"github.com/dmitry-davydov/go-events-state/internal/models"
	"github.com/labstack/echo"
)

func list() func (c echo.Context) error {
	return func (c echo.Context) error {

	}
}
func getOne(models.EventId) func (c echo.Context) error {
	return func (c echo.Context) error {
		
	}
}
func delete(models.EventId eventID) func (c echo.Context) error {
	return func (c echo.Context) error {
		
	}
}
func create() func (c echo.Context) error {
	return func (c echo.Context) error {
		ev := new(models.Event)	
		ev.GenerateEventId()
		ev.Status = models.EventStatusCreated
	}
}
func update(eventId models.EventId, event request.Event) func (c echo.Context) error {
	return func (c echo.Context) error {
		
	}
}
