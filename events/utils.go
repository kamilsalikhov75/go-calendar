package events

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

func getNextId() string {
	return uuid.New().String()
}

func ParseDateStr(dateStr string) (time.Time, error) {
	t, err := dateparse.ParseLocal(dateStr)
	if err != nil {
		return time.Time{}, errors.New(notValidDateError)
	}

	return t.In(time.Local), nil
}
