package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/handlers"
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	port := os.Getenv("PORT")
	r := chi.NewMux()
	db := "dbstore"
	passwordHash := "hash"

	userStore := dbstore.NewUserStore(dbstore.NewUserStoreParams{
		DB:           db,
		PasswordHash: passwordHash,
	})

	r.Handle("/*", public())

	r.Get("/", handlers.Make(handlers.HandleHome))
	r.Get("/login", handlers.Make(handlers.HandleLogin))
	r.Get("/signup", handlers.Make(handlers.HandleSignup))
	r.Get("/complete", handlers.Make(handlers.HandleFinished))
	r.Get("/processing", handlers.Make(handlers.HandleProcessing))
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println("Here in post login")
		fmt.Print("Email:  ", email, " Password: ", password)
	})
	r.Post("/signup", handlers.NewPostSignupHandler(
		handlers.PostSignupHandlerParams{UserStore: userStore}))
	http.ListenAndServe(port, r)
}

func public() http.Handler {
	return http.StripPrefix("/public", http.FileServer(http.Dir("public")))
}
