package storage

import (
	"context"
	"sync"
	"time"

	"github.com/dmitry-davydov/go-events-state/internal/models"
)

// Storage хранит все данные о событиях
type Storage struct {
	mut  sync.RWMutex
	data map[string]*models.Event

	interruptedValidPeriod int

	exitContext context.Context
}

func (t *Storage) finishSateRunner() {
	ticker := time.NewTicker(time.Second * time.Duration(t.interruptedValidPeriod))

	for {
		select {
		case <-t.exitContext.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			t.changeStates()
		}

	}
}

func (t *Storage) changeStates() {
	t.mut.Lock()
	for _, event := range t.data {
		if !event.CanFinish(t.interruptedValidPeriod) {
			continue
		}

		event.SetFinishState()
	}

	t.mut.Unlock()
}

// Delete удаляет *models.Event по string
func (t *Storage) Delete(ID string) bool {
	t.mut.Lock()
	defer t.mut.Unlock()

	if _, ok := t.data[ID]; !ok {
		return false
	}

	delete(t.data, ID)

	return true
}

// SetEvent заменит или добавит в сторейдж *models.Event структуру по string
func (t *Storage) SetEvent(event *models.Event) {
	t.mut.Lock()
	defer t.mut.Unlock()

	t.data[event.ID] = event
}

// Event вернет *models.Event или nil если ее нет в сторейдже
func (t *Storage) Event(ID string) *models.Event {
	t.mut.RLock()
	defer t.mut.RUnlock()

	event, ok := t.data[ID]
	if !ok {
		return nil
	}

	return event
}

func (t *Storage) All() []*models.Event {
	var all []*models.Event

	t.mut.RLock()
	defer t.mut.RUnlock()

	for _, e := range t.data {
		all = append(all, e)
	}

	return all
}

// NewStorage это "конструктор" для создания Storage
func NewStorage(interruptedValidPeriod int, context context.Context) *Storage {
	t := new(Storage)
	t.interruptedValidPeriod = interruptedValidPeriod
	t.data = make(map[string]*models.Event)
	t.exitContext = context

	go t.finishSateRunner()

	return t
}
