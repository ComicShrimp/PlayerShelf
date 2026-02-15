// Package internal handles all the connection between the layers in the arch
package internal

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ComicShrimp/PlayerShelf/internal/infra/database"
	"github.com/joho/godotenv"
)

func Run(ctx context.Context) {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create DB connection
	dbPool, err := database.NewPostgresPool(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to DB %v\n", err)
	}

	defer dbPool.Close()

	// Test query
	var greeting string
	err = dbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
