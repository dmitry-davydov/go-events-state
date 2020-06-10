package response

import (
	"time"

	"github.com/dmitry-davydov/go-events-state/internal/models"
)

// Event структура для http api
// {
//   "type": "stream",
//   "id": "63efedc2-7f03-4760-87be-31f30f38526d",
//   "attributes": {
//     "created": "2017-07-18T10:20:34.478621543Z",
//     "state": "created"
//   }
// }
type Event struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CreatedAt time.Time          `json:"created"`
		State     models.EventStatus `json:"state"`
	} `json:"attributes"`
}

func (t *Event) Map(event *models.Event) {
	t.Type = "stream"
	t.ID = event.ID
	t.Attributes.CreatedAt = event.CreatedAt
	t.Attributes.State = event.Status
}
