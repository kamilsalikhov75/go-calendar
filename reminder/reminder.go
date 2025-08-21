package reminder

import (
	"time"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
	timer   *time.Timer
	notify  func(string)
}

func NewRemider(message string, at time.Time, notify func(string)) *Reminder {
	return &Reminder{
		Message: message,
		At:      at,
		Sent:    false,
		notify:  notify,
	}
}

func (r *Reminder) Start() {
	timer := time.AfterFunc(r.At.Sub(time.Now()), r.Send)

	r.timer = timer
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	r.notify(r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	if r.timer != nil {
		r.timer.Stop()
	}
}
