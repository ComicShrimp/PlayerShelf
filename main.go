package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ComicShrimp/PlayerShelf/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	internal.Run(ctx)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	err := r.Run()
	if err != nil {
		log.Fatalf("Gin server failed to run: %v", err)
	}
}
