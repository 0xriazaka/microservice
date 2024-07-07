package app

import (
	"net/http"
)

type App struct {
	router http.Handler
}
	
func New() *App {
	app := &App {
		router: LoadRoutes()
	}
	return app
}