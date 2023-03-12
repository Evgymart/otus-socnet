package core

import (
	"otus/socnet/db"
	"otus/socnet/models"
)

func AddUser(user *models.User) Message {
	database := db.GetDatabase()
	err := db.AddUser(database, user)
	if err != nil {
		return Message{
			Status:  "Error",
			Message: err.Error(),
		}
	}

	return Message{
		Status:  "OK",
		Message: "",
	}
}
