package core

import (
	"otus/socnet/db"
	"otus/socnet/models"
)

type TokenMessage struct {
	Status  string
	Message string
	Token   string
}

func Login(creds *models.Credentials) TokenMessage {
	database := db.GetDatabase()
	storedPassword, err := db.Login(database, creds)
	if err != nil {
		return TokenMessage{
			Status:  "Error",
			Message: err.Error(),
			Token:   "",
		}
	}

	response := TokenMessage{
		Status:  "OK",
		Message: "",
		Token:   "",
	}

	success := compareHash(creds.Password, storedPassword)
	if success {
		response.Token = "TOKEN"
	} else {
		response.Message = "Bad credentials"
	}

	return response
}
