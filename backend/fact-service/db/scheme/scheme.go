package scheme

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SetupScheme(ctx context.Context, conn *pgx.Conn) error {
	schemeQuery := `CREATE TABLE IF NOT EXISTS facts (
        id SERIAL PRIMARY KEY,
        text TEXT NOT NULL
    )`
	_, err := conn.Exec(ctx, schemeQuery)
	if err != nil {
		return fmt.Errorf("failed to create scheme: %w", err)
	}
	return nil
}
