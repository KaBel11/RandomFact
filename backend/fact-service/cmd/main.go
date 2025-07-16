package main

import (
	"context"
	"fmt"

	"github.com/KaBel11/RandomFact/fact-service/internal/api"
	"github.com/KaBel11/RandomFact/fact-service/internal/repository"
	"github.com/KaBel11/RandomFact/fact-service/internal/service"
)

func main() {
	repo := repository.NewFactsRepository()
    svc := service.NewFactsService(repo)
    h := api.NewFactsHandler(svc)
	api := api.New(h)

	err := api.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start api:", err)
	}
}