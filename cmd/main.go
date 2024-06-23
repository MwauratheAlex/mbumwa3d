package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/handlers"
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	port := os.Getenv("PORT")
	r := chi.NewMux()

	r.Handle("/*", public())

	r.Get("/", handlers.Make(handlers.HandleHome))
	r.Get("/complete", handlers.Make(handlers.HandleFinished))
	r.Get("/processing", handlers.Make(handlers.HandleProcessing))
	http.ListenAndServe(port, r)
}

func public() http.Handler {
	return http.StripPrefix("/public", http.FileServer(http.Dir("public")))
}
