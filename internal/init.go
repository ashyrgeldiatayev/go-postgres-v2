package internal

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewConnection(config *Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User, 
		config.Password, 
		config.Host, 
		config.Port, 
		config.DBName, 
		config.SSLMode)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
