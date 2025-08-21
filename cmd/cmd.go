package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/kamilsalikhov75/go-calendar/calendar"
	"github.com/kamilsalikhov75/go-calendar/events"
	"github.com/kamilsalikhov75/go-calendar/logger"
	"github.com/kamilsalikhov75/go-calendar/storage"
)

type Cmd struct {
	calendar   *calendar.Calendar
	logs       []Log
	logStorage storage.Store
	logger     *logger.Logger
}

func NewCmd(c *calendar.Calendar) *Cmd {
	logStorage := storage.NewJsonStorage("calendar.log.json")
	logger, err := logger.NewLogger("app.log")

	if err != nil {
		fmt.Println(err)
	}

	return &Cmd{
		calendar:   c,
		logStorage: logStorage,
		logger:     logger,
	}
}

func (c *Cmd) Run() {
	defer func() {
		err := c.calendar.Save()
		if err != nil {
			text := calendar.SaveError + " " + err.Error()
			fmt.Println(text)
			c.AddLog(text)
		}
	}()
	err := c.calendar.Load()
	if err != nil {
		text := loadError + err.Error()
		fmt.Println(text)
		c.AddLog(text)
	}

	err = c.LoadLogs()
	if err != nil {
		text := loadLogsError + err.Error()
		fmt.Println(text)
		c.AddLog(text)
	}

	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	go func() {
		for msg := range c.calendar.Notification {
			fmt.Println(msg)
			c.AddLog(msg)
		}
	}()
	p.Run()

}

func (c *Cmd) executor(input string) {
	c.AddLog(input)
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println(err)
		c.AddLog(err.Error())
		return
	}
	cmd := strings.ToLower(parts[0])
	switch cmd {
	case CommandAdd:
		if len(parts) < 4 {
			fmt.Println(addFormatWarning)
			c.AddLog(addFormatWarning)
			return
		}

		title := parts[1]
		date := parts[2]
		priority := events.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		c.logger.Info("c.calendar.AddEvent(" + title + "," + date + "," + string(priority) + ")")
		if err != nil {
			fmt.Println(err)
			c.logger.Error(err.Error())
			c.AddLog(err.Error())
		} else {
			text := "Событие: " + e.Title + " добавлено"
			fmt.Println(text)
			c.logger.Info(text)
			c.AddLog(text)
		}
	case CommandList:
		events, err := c.calendar.GetEvents()
		if err != nil {
			fmt.Println(err)
			c.AddLog(err.Error())
		}

		text := "События:"
		for _, e := range events {
			text = text + "\n" + "--ID: " + e.ID + "\n  Название: " + e.Title + "\n  Дата: " + e.StartAt.Format("2 January 2006 15:04") + "\n  Приоритет: " + e.Priority.Translate()
		}

		fmt.Println(text)
		c.AddLog(text)
	case CommandRemove:
		if len(parts) < 2 {
			fmt.Println(removeFormatWarning)
			c.AddLog(removeFormatWarning)
			return
		}

		id := parts[1]

		err := c.calendar.DeleteEvent(id)
		c.logger.Info("c.calendar.DeleteEvent(" + id + ")")
		if err != nil {
			fmt.Println(err)
			c.AddLog(err.Error())
		} else {
			text := "Событие с id " + id + " удалено"
			fmt.Println(text)
			c.AddLog(text)
		}
	case CommandUpdate:
		if len(parts) < 5 {
			fmt.Println(updateFormatWarning)
			c.AddLog(updateFormatWarning)
			return
		}

		id := parts[1]
		title := parts[2]
		date := parts[3]
		priority := events.Priority(parts[4])

		err := c.calendar.EditEvent(id, title, date, priority)
		if err != nil {
			fmt.Println(err)
			c.AddLog(err.Error())
		} else {
			text := "Событие: " + title + "обновлено"
			c.logger.Info(text)
			fmt.Println(text)
			c.AddLog(text)
		}
	case CommandHelp:
		fmt.Println("Список команд:")
		fmt.Println("---", CommandAdd, "- добавить событие")
		fmt.Println("---", CommandList, "- показать все события")
		fmt.Println("---", CommandRemove, "- удалить событие")
		fmt.Println("---", CommandUpdate, "- обновить событие")
		fmt.Println("---", CommandHelp, "- показать справку")
		fmt.Println("---", CommandRemind, `- напомнить о событии за "1 2 3" "час минута секунда"`)
		fmt.Println("---", CommandStop, "- отменить напоминание о событии")
		fmt.Println("---", CommandExit, "- выход из программы")
	case CommandRemind:
		if len(parts) < 3 {
			fmt.Println(remindFormatWarning)
			c.AddLog(remindFormatWarning)
			return
		}

		id := parts[1]
		duration := parts[2]

		err := c.calendar.SetEventReminder(id, duration)
		if err != nil {
			fmt.Println(err)
			c.AddLog(err.Error())
			return
		}
	case CommandStop:
		if len(parts) < 2 {
			fmt.Println(stopFormatWarning)
			c.AddLog(stopFormatWarning)
			return
		}

		id := parts[1]

		err := c.calendar.CancelEventReminder(id)
		if err != nil {
			fmt.Println(err)
			c.AddLog(err.Error())
			return
		}
	case CommandExit:
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(calendar.SaveError, err)
			c.AddLog(calendar.SaveError, err.Error())
		}
		err = c.SaveLogs()
		if err != nil {
			fmt.Println(logsSaveError, err)
		}

		close(c.calendar.Notification)
		c.logger.Close()
		os.Exit(0)
	case CommandLog:
		c.PrintLogs()
	default:
		text := "Неизвестная команда:\nВведите 'help' для списка команд"
		c.logger.Error(text)
		fmt.Println(text)
		c.AddLog(text)
	}

}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	userInput := d.CurrentLine()
	suggestions := []prompt.Suggest{
		{Text: CommandAdd, Description: "Добавить событие"},
		{Text: CommandList, Description: "Показать все события"},
		{Text: CommandRemove, Description: "Удалить событие"},
		{Text: CommandUpdate, Description: "Обновить событие"},
		{Text: CommandHelp, Description: "Показать справку"},
		{Text: CommandRemind, Description: `Напомнить о событии за "1 2 3" "час минута секунда"`},
		{Text: CommandStop, Description: "Отменить напоминание о событии"},
		{Text: CommandLog, Description: "Показать логи"},
		{Text: CommandExit, Description: "Выйти из программы"},
	}

	events, err := c.calendar.GetEvents()
	eventSuggestions := []prompt.Suggest{}
	if err == nil {
		for _, e := range events {
			if !strings.Contains(userInput, e.Title) {
				continue
			}
			eventSuggestions = append(eventSuggestions, prompt.Suggest{
				Text:        e.ID,
				Description: "Событие: " + e.Title,
			})

		}
	}

	parts, err := shlex.Split(userInput)
	command := ""
	if err == nil {
		if len(parts) > 0 {
			command = strings.ToLower(parts[0])
		}
	}

	filtredSuggestions := []prompt.Suggest{}

	for _, s := range suggestions {

		if command == "" {
			filtredSuggestions = append(filtredSuggestions, s)
			continue
		}
		if strings.Contains(s.Text, command) {
			filtredSuggestions = append(filtredSuggestions, s)
		}
	}

	return append(filtredSuggestions, eventSuggestions...)

}
