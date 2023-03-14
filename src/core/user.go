package core

import (
	"otus/socnet/db"
	"otus/socnet/models"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(user *models.User) Message {
	hashPassword(user)
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

func hashPassword(user *models.User) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		panic("Could not hash a password")
	}
	user.Password = string(hashed)
}
