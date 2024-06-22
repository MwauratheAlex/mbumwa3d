package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/handlers"
)

func main() {
	r := chi.NewMux()

	r.Handle("/*", public())

	r.Get("/", handlers.Make(handlers.HandleHome))
	http.ListenAndServe(":3000", r)
}

func public() http.Handler {
	return http.StripPrefix("/public", http.FileServer(http.Dir("public")))
}
