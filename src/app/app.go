package app

import (
	"net/http"
	"otus/socnet/api"
)

type App struct {
	Mux *http.ServeMux
}

var (
	appInstance *App = nil
)

func NewApp() (*App, error) {
	mux := http.NewServeMux()

	return &App{
		Mux: mux,
	}, nil
}

func (app *App) Start() {
	api.InitApi(app.Mux)
	http.ListenAndServe("app:80", app.Mux)
	appInstance = app
}

func GetApp() *App {
	if appInstance == nil {
		panic("App is not initiated")
	}

	return appInstance
}
