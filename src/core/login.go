package core

import (
	"otus/socnet/db"
	"otus/socnet/models"
)

type TokenMessage struct {
	Status  string `json:"token"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Login(creds *models.Credentials) (bool, int) {
	database := db.GetReadDb()
	id, storedPassword, _ := db.Login(database, creds)
	return compareHash(creds.Password, storedPassword), id
}
