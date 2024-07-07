package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// handler
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	// router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/welcome", basicHandler)

	// http server
	s := &http.Server{
		Addr:    ":7070",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}
