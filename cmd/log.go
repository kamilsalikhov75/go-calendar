package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Log struct {
	At  time.Time
	Log string
}

var mu sync.Mutex

func (c *Cmd) AddLog(log ...string) {
	mu.Lock()
	c.logs = append(c.logs, Log{At: time.Now(), Log: strings.Join(log, "")})
	mu.Unlock()
}

func (c *Cmd) PrintLogs() {
	for _, l := range c.logs {
		fmt.Println(l.At.Format("2 January 2006 15:04:05"), l.Log)
	}
}

func (c *Cmd) SaveLogs() error {
	data, err := json.Marshal(c.logs)

	if err != nil {
		return err
	}

	err = c.logStorage.Save(data)
	return err
}

func (c *Cmd) LoadLogs() error {
	data, err := c.logStorage.Load()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c.logs)
	return err
}
