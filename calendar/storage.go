package calendar

import (
	"encoding/json"
)

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)

	if err != nil {
		return err
	}

	err = c.storage.Save(data)
	return err
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c.calendarEvents)
	return err
}
