package core

import (
	"otus/socnet/db"
	"otus/socnet/models"
)

type Message struct {
	Status  string
	Message string
}

func AddUser(user *models.User) Message {
	user.Password = hashPassword(user.Password)
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
