package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

type EventStatus uint8
type EventType uint8

const (
	EventStatusCreated EventStatus = iota
	EventStatusActive
	EventStatusInterrupted
	EventStatusFinished
	EventStatusUnknown EventStatus = 100
)

const (
	EventTypeStream EventType = iota
)

func (t *EventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "stream":
		*t = EventTypeStream
	}

	return nil
}

func (t EventType) MarshalJSON() ([]byte, error) {

	var s string

	switch t {
	case EventTypeStream:
		s = "stream"
	}

	return json.Marshal(s)
}

func (t *EventStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	default:
		*t = EventStatusUnknown
	case "created":
		*t = EventStatusCreated
	case "active":
		*t = EventStatusActive
	case "interrupted":
		*t = EventStatusInterrupted
	case "finished":
		*t = EventStatusFinished
	}

	return nil
}

func (t EventStatus) MarshalJSON() ([]byte, error) {

	var s string

	switch t {
	case EventStatusCreated:
		s = "created"
	case EventStatusActive:
		s = "active"
	case EventStatusInterrupted:
		s = "interrupted"
	case EventStatusFinished:
		s = "finished"
	}

	return json.Marshal(s)
}

type Event struct {
	ID             string
	Status         EventStatus
	Type           EventType
	CreatedAt      time.Time
	StatusUpdatedAt time.Time
}

func (t *Event) CanFinish(interruptedValidPeriod int) bool {

	return t.Status == EventStatusInterrupted && time.Now().Second() - t.StatusUpdatedAt.Second() >= interruptedValidPeriod
}

func (t *Event) SetFinishState() {
	t.Status = EventStatusFinished
}

func (t *Event) SetState(newState EventStatus) {
	allowedStatuses := map[EventStatus]map[EventStatus]struct{}{
		EventStatusCreated: {
			EventStatusActive: {},
		},
		EventStatusActive: {
			EventStatusInterrupted: {},
		},
		EventStatusInterrupted: {
			EventStatusActive: {},
		},
	}

	if sub, ok := allowedStatuses[t.Status]; ok {
		if _, ok := sub[newState]; ok {
			t.Status = newState
			t.StatusUpdatedAt = time.Now()
		}
	}
}

func generateEventId() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return id.String()
}

func NewEvent() *Event {
	event := new(Event)
	event.Status = EventStatusCreated
	event.ID = generateEventId()
	event.CreatedAt = time.Now()
	event.Type = EventTypeStream

	return event
}
