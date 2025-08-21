package main

import (
	"github.com/kamilsalikhov75/go-calendar/calendar"
	"github.com/kamilsalikhov75/go-calendar/cmd"
	"github.com/kamilsalikhov75/go-calendar/storage"
)

func main() {
	s := storage.NewJsonStorage("calendar.json")
	c := calendar.NewCalendar(s)

	cli := cmd.NewCmd(c)
	cli.Run()

}
