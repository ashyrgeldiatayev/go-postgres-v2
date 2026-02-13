package internal

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

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
