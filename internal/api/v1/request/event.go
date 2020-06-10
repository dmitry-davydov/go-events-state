package request

import "github.com/dmitry-davydov/go-events-state/internal/models"

type Event struct {
	State models.EventStatus `json:"state"`
}

// IsValid проверят на валидность запроса от клиента
func (t Event) IsValid() bool {
	return t.State != models.EventStatusUnknown
}
