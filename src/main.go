package main

import (
	"otus/socnet/app"
	"otus/socnet/db"
)

func main() {
	db.InitDatabase()
	app, err := app.NewApp()
	if err != nil {
		panic(err.Error())
		return
	}

	app.Start()
}
