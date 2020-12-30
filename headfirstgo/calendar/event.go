package calendar

import (
	"errors"
	"unicode/utf8"
)

// Event is a struct with Title(string), Date(another struct)
type Event struct {
	title string
	Date
}

// Title return title for Event
func (e *Event) Title() string {
	return e.title
}

// SetTitle sets title for Event
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title lenth:1-30")
	}

	e.title = title
	return nil
}
