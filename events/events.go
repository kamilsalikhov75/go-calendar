package events

import (
	"errors"
	"time"

	"github.com/google/shlex"
	"github.com/kamilsalikhov75/go-calendar/reminder"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

func NewEvent(title string, dateStr string, priority Priority) (*Event, error) {
	err := ValidateTitle(title)
	if err != nil {
		return nil, err
	}

	t, err := ParseDateStr(dateStr)
	if err != nil {
		return nil, err
	}

	err = priority.Validate()
	if err != nil {
		return nil, err
	}

	event := Event{
		ID:       getNextId(),
		Title:    title,
		StartAt:  t,
		Priority: priority,
		Reminder: nil,
	}
	return &event, nil
}

func (e *Event) Update(title string, dateStr string, priority Priority) error {
	if title != "" {
		err := ValidateTitle(title)
		if err != nil {
			return err
		}
		e.Title = title
	}

	if dateStr != "" {
		t, err := ParseDateStr(dateStr)
		if err != nil {
			return err
		}
		e.StartAt = t
	}

	if priority != "" {
		err := priority.Validate()
		if err != nil {
			return err
		}
		e.Priority = priority
	}

	return nil
}

func (e *Event) AddReminder(durationPartsString string, notify func(string)) error {
	durationParts, err := shlex.Split(durationPartsString)
	if err != nil {
		return err
	}
	if len(durationParts) != 3 {
		return errors.New(notValidDurationError)
	}

	durationString := durationParts[0] + "h" + durationParts[1] + "m" + durationParts[2] + "s"
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}
	at := e.StartAt.Add(time.Duration(-duration))

	message := e.Title + " через " + durationString

	e.Reminder = reminder.NewRemider(message, at, notify)
	e.Reminder.Start()
	return nil
}

func (e *Event) RemoveReminder() {
	if e.Reminder != nil {
		e.Reminder.Stop()
		e.Reminder = nil
	}
}
