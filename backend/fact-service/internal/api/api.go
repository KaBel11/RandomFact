package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KaBel11/RandomFact/fact-service/internal/service"
)

type Api struct {
	router http.Handler
}

func New(service *service.FactsService) *Api {
	handler := NewFactsHandler(service)
	api := &Api{
		router: loadRoutes(handler),
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