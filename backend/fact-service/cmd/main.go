package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KaBel11/RandomFact/fact-service/db"
	"github.com/KaBel11/RandomFact/fact-service/db/scheme"
	"github.com/KaBel11/RandomFact/fact-service/internal/api/handler"
	"github.com/KaBel11/RandomFact/fact-service/internal/repository"
	"github.com/KaBel11/RandomFact/fact-service/internal/service"
	"github.com/KaBel11/RandomFact/fact-service/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not initialize .env filed: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	dbConn, err := db.NewDatabase(ctx)
	if err != nil {
		log.Fatalf("Could not initialize DB connection: %s", err)
	}
	defer dbConn.Close(ctx)

	if err := dbConn.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	err = scheme.SetupScheme(ctx, dbConn)
	if err != nil {
		log.Fatalf("Could not setup DB scheme: %s", err)
	}

	repository := repository.NewFactRepository(dbConn)
	service := service.NewFactService(repository)
	handler := handler.NewFactHandler(service)
	router := router.SetupRouter(handler)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}
