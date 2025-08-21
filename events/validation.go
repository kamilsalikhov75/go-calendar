package events

import (
	"errors"
	"regexp"
)

func isValidTitle(title string) bool {
	pattern := `^[a-zA-Zа-яА-Я0-9 ,\.]{3,100}$`
	matched, _ := regexp.MatchString(pattern, title)
	return matched
}

func ValidateTitle(title string) error {
	if isValidTitle(title) {
		return nil
	} else {
		return errors.New("Неверный формат названия")
	}
}
