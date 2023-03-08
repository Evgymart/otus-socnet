package main

import "otus/socnet/app"

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic(err.Error())
		return
	}

	app.Start()
}
