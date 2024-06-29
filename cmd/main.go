package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/handlers"
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")
	r := chi.NewMux()

	userStore := dbstore.NewUserStore()

	authMiddleware := middleware.NewAuthMiddleware("Authorization", userStore)

	r.Group(func(r chi.Router) {
		r.Use(
			//middleware.TextHTMLMiddleware,
			authMiddleware.AddUserToContext,
		)

		r.Handle("/*", public())

		r.Get("/", handlers.Make(handlers.HandleHome))
		r.Get("/login", handlers.Make(handlers.HandleLogin))
		r.Get("/signup", handlers.Make(handlers.HandleSignup))
		r.Get("/complete", handlers.Make(handlers.HandleFinished))
		r.Get("/processing", handlers.Make(handlers.HandleProcessing))
		r.Get("/usermenu", handlers.Make(handlers.GetUserMenu))
		r.Post("/login", handlers.NewPostLoginHandler(
			handlers.PostLoginHandlerParams{UserStore: userStore},
		))
		r.Post("/signup", handlers.NewPostSignupHandler(
			handlers.PostSignupHandlerParams{UserStore: userStore},
		))
		r.Post("/logout", handlers.Make(handlers.PostLogout))
		r.Post("/print", handlers.Make(handlers.PostPrint))
	})
	http.ListenAndServe(port, r)
}

func public() http.Handler {
	return http.StripPrefix("/public", http.FileServer(http.Dir("public")))
}
