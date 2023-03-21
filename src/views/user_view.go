package views

import (
	"strings"
	"time"
)

type Date time.Time

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Birthdate Date   `json:"birthdate`
	Gender    string `json:"gender"`
	Biography string `json:"biography"`
	City      string `json:"city"`
}

const DateFormat = "02.01.2006"

func (date *Date) UnmarshalJSON(bytes []byte) error {
	stringValue := strings.Trim(string(bytes), `"`)
	if stringValue == "" {
		return nil
	}

	time, err := time.Parse(DateFormat, stringValue)
	if err != nil {
		return err
	}

	*date = Date(time)
	return nil
}
