package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KaBel11/RandomFact/fact-service/internal/api"
	"github.com/KaBel11/RandomFact/fact-service/internal/repository"
	"github.com/KaBel11/RandomFact/fact-service/internal/service"
	"github.com/jackc/pgx/v5"
)

func Start(ctx context.Context) error {
	databaseURL := "postgres://postgres:admin@localhost:5432/randomfact_db?sslmode=disable"

	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			fmt.Println("failed to close database connection", err)
		}
	}()

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS facts (
        id SERIAL PRIMARY KEY,
        text TEXT NOT NULL
    )`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	ch := make(chan error, 1)

	repository := repository.NewPostgresFactsRepository(conn)
	service := service.NewFactsService(repository)
	router := api.NewFactsHandler(service)

	server := &http.Server{
		Addr:    ":3000",
		Handler: api.LoadRoutes(router),
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
