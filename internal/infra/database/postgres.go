// Package database handles all the concrete DB connection
package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	dbConfig := getPostgresConfig(connectionString)

	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
	}

	return conn, err
}

func RunMigrations(connectionString string) error {
	if os.Getenv("RUN_MIGRATIONS") != "true" {
		log.Println("Automatic DB migrations are disabled!")
		return nil
	}

	// TODO: Refactor
	convertedString := connectionString
	if strings.HasPrefix(convertedString, "postgresql://") {
		convertedString = strings.Replace(convertedString, "postgresql://", "pgx5://", 1)
	} else if strings.HasPrefix(convertedString, "postgres://") {
		convertedString = strings.Replace(convertedString, "postgres://", "pgx5://", 1)
	}

	m, err := migrate.New(
		"file://internal/infra/database/migrations",
		convertedString,
	)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	log.Println("Running database migrations...")
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("No new migrations to apply.")
			return nil
		}
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Migrations completed successfully!")
	return nil
}

func getPostgresConfig(connectionString string) *pgxpool.Config {
	dbConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse Config for Postgres %v\n", err)
	}

	return dbConfig
}
