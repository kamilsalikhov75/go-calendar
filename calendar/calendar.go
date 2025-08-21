package calendar

import (
	"errors"

	"github.com/kamilsalikhov75/go-calendar/events"
	"github.com/kamilsalikhov75/go-calendar/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
}

func NewCalendar(s storage.Store) *Calendar {

	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        s,
		Notification:   make(chan string),
	}
}

func (c Calendar) GetEvents() (map[string]*events.Event, error) {

	if len(c.calendarEvents) == 0 {
		return nil, errors.New(noEventsError)
	}
	return c.calendarEvents, nil

}

func (c *Calendar) AddEvent(title string, date string, priority events.Priority) (*events.Event, error) {

	e, err := events.NewEvent(title, date, priority)

	if err != nil {
		return nil, errors.New(addError + err.Error())
	}

	c.calendarEvents[e.ID] = e

	return e, nil
}

func (c *Calendar) DeleteEvent(id string) error {

	_, ok := c.calendarEvents[id]

	if !ok {
		return errors.New(deleteError + notFoundError)
	}

	delete(c.calendarEvents, id)

	return nil
}

func (c *Calendar) EditEvent(id string, title string, date string, priority events.Priority) error {

	e, ok := c.calendarEvents[id]
	if !ok {
		return errors.New(editError + notFoundError)
	}
	err := e.Update(title, date, priority)
	if err != nil {
		return errors.New(editError + err.Error())
	}

	return nil
}
