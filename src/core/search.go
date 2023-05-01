package core

import (
	"otus/socnet/db"
	"otus/socnet/views"
)

type SearchData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func SearchUsers(firstName string, lastName string) ([]views.User, error) {
	database := db.GetReadDb()
	return db.SearchUsers(database, firstName, lastName)
}
