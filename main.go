package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ComicShrimp/PlayerShelf/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	ctx := context.Background()

	internal.Run(ctx)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user := struct {
			Message string `json:"message"`
		}{
			Message: "Hello World!",
		}

		// This sets Content-Type to application/json automatically
		render.JSON(w, r, user)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Unable to run Chi server %v\n", err)
	}
}
