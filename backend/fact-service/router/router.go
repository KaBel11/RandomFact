package router

import (
	"net/http"

	"github.com/KaBel11/RandomFact/fact-service/internal/api/handler"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(factH *handler.FactHandler) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/api/facts", func(r chi.Router) {
		r.Get("/", factH.List)
		r.Get("/{id}", factH.GetById)
		r.Get("/random", factH.GetRandomFact)
		r.Post("/", factH.Create)
		r.Put("/{id}", factH.Update)
		r.Delete("/{id}", factH.Delete)
	})

	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	return router
}