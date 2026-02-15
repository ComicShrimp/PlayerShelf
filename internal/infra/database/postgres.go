// Package database handles all the concrete DB connection
package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	dbConfig := getPGConfig(connectionString)

	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
	}

	return conn, err
}

func getPGConfig(connectionString string) *pgxpool.Config {
	dbConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse Config for Postgres %v\n", err)
	}

	// TODO: Ativar isso somente para o localhost
	dbConfig.ConnConfig.TLSConfig.InsecureSkipVerify = true

	return dbConfig
}
