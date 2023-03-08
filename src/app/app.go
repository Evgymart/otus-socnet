package app

import (
	"net/http"
	"otus/socnet/api"
)

type App struct {
	Mux *http.ServeMux
}

func NewApp() (*App, error) {
	mux := http.NewServeMux()
	return &App{
		Mux: mux,
	}, nil
}

func (app *App) Start() {
	api.InitApi(app.Mux)
	http.ListenAndServe("app:80", app.Mux)
}
