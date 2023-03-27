package core

import (
	"otus/socnet/db"
	"otus/socnet/views"
)

type SearchData struct {
	Name string `json:"name"`
}

func SearchUsers(name string) ([]views.User, error) {
	database := db.GetDatabase()
	return db.SearchUsers(database, name)
}
