package core

import (
	"errors"
	"otus/socnet/db"
)

type SwapUserNamesData struct {
	FirstId int `json:"first_id"`
	LastId  int `json:"last_id"`
}

func SwapUserNames(data *SwapUserNamesData) error {
	if data == nil {
		return errors.New("Bad data provided")
	}

	database := db.GetWriteDb()
	return db.SwapUserNames(database, data.FirstId, data.LastId)
}
