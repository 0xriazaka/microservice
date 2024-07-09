package app

import (
	"api/register"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes
func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/registration", loadRegisterRoutes)

	return router

}
func loadRegisterRoutes(router chi.Router) {
	register := &register.Registration{}
	router.Post("/", register.Create)
	router.Get("/", register.List)
	router.Get("/{id}", register.GetByID)
	router.Put("/{id}", register.UpdateByID)
	router.Delete("/{id}", register.DeleteByID)
}
