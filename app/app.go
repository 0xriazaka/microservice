package app

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

// router
func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	// http server
	s := &http.Server{
		Addr:    ":7070",
		Handler: a.router,
	}
	err := s.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	} else {
		return nil
	}

}
