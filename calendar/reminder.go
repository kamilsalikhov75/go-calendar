package calendar

import (
	"errors"
)

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) SetEventReminder(id string, duration string) error {
	e, ok := c.calendarEvents[id]
	if !ok {
		return errors.New(notFoundError)
	}
	err := e.AddReminder(duration, c.Notify)
	if err != nil {
		return err
	}

	return err
}

func (c *Calendar) CancelEventReminder(id string) error {
	e, ok := c.calendarEvents[id]

	if !ok {
		return errors.New(notFoundError)
	}

	e.RemoveReminder()

	return nil
}
