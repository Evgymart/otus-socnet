package core

import (
	"otus/socnet/db"
	"otus/socnet/models"
)

func Register(user *models.User) Message {
	user.Password = hashPassword(user.Password)
	database := db.GetDatabase()
	err := db.AddUser(database, user)
	if err != nil {
		return ResponseError(err)
	}

	return ResponseOK()
}
