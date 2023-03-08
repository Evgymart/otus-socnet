package models

import (
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Birthdate time.Time
	Gender    string
	Biography string
	City      string
}
