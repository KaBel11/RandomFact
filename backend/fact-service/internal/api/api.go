package api

import (
	"context"
	"fmt"
	"net/http"
)

type Api struct {
	router http.Handler
}

func New(factsHandler *FactsHandler) *Api {
	api := &Api{
		router: loadRoutes(factsHandler),
	}
	return api
}

func (a *Api) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to listen to server: %w", err)
	}

	return nil
}