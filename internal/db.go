package internal

import (
	"context"

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

func MigrateBooks(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), CreateBooksTableQuery)
	if err != nil {
		return err
	}

	return nil
}
