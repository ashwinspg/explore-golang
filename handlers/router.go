package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//GetRouter - router configuration
func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/ping", PingHandler)
	router.Get("/movies/{id}", GetMovieHandler)

	return router
}
