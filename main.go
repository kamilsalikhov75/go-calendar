package main

import (
	"github.com/kamilsalikhov75/app/calendar"
	"github.com/kamilsalikhov75/app/cmd"
	"github.com/kamilsalikhov75/app/storage"
)

func main() {
	s := storage.NewJsonStorage("calendar.json")
	c := calendar.NewCalendar(s)

	cli := cmd.NewCmd(c)
	cli.Run()

}
