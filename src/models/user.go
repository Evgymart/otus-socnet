package models

import (
	"encoding/json"
	"strings"
	"time"
)

type Date time.Time

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Birthdate Date   `json:"birthdate`
	Gender    string `json:"gender"`
	Biography string `json:"biography"`
	City      string `json:"city"`
}

const DateFormat = "02.01.2006"

func BuildUser(userData []byte) (*User, error) {
	var user User
	err := json.Unmarshal(userData, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

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

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(date).Format(DateFormat) + `"`), nil
}

func (date Date) ToString() string {
	return time.Time(date).Format(DateFormat)
}
