package main

import (
	"context"
	"fmt"

	"github.com/KaBel11/RandomFact/fact-service/internal/api"
	"github.com/KaBel11/RandomFact/fact-service/internal/repository"
	"github.com/KaBel11/RandomFact/fact-service/internal/service"
)

func main() {
	repository := repository.NewFactsRepository()
    service := service.NewFactsService(repository)
	api := api.New(service)

	err := api.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start api:", err)
	}
}