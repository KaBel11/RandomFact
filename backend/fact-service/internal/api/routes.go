package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LoadRoutes(factsHandler *FactsHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/facts", func(r chi.Router) {
		r.Post("/", factsHandler.Create)
		r.Get("/", factsHandler.List)
		r.Get("/{id}", factsHandler.GetById)
		r.Get("/random", factsHandler.GetRandomFact)
		r.Put("/{id}", factsHandler.Update)
		r.Delete("/{id}", factsHandler.Delete)
	})

	return router
}