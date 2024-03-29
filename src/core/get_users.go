package core

import (
	"otus/socnet/db"
	"otus/socnet/views"
)

type GetUserData struct {
	Email string `json:"email"`
}

func GetUsers(limit int) ([]views.User, error) {
	database := db.GetReadDb()
	return db.GetUsers(database, limit)
}

func GetUser(email string) (*views.User, error) {
	database := db.GetReadDb()
	return db.GetUser(database, email)
}
