package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func MigrateBooks(conn *pgx.Conn) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		author VARCHAR(255),
		title VARCHAR(255),
		publisher VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
