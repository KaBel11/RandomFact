package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/KaBel11/RandomFact/fact-service/internal/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := application.Start(ctx)
	if err != nil {
		fmt.Println("failed to start application:", err)
	}
}
