package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/handlers"
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

var Environment string = "dev"

func init() {
	switch Environment {
	case "docker-prod":
		initializers.LoadEnvVariables()
		os.Setenv("env", "production")
		break
	case "dev":
		initializers.LoadEnvVariables()
	}
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
		r.Get("/dashboard", handlers.Make(handlers.HandleDashboard))

		r.Post("/login", handlers.NewPostLoginHandler(
			handlers.PostLoginHandlerParams{UserStore: userStore},
		))
		r.Post("/signup", handlers.NewPostSignupHandler(
			handlers.PostSignupHandlerParams{UserStore: userStore},
		))
		r.Post("/logout", handlers.Make(handlers.PostLogout))
		r.Post("/print", handlers.Make(handlers.PostPrint))
		r.Post("/payment", handlers.Make(handlers.PostPayment))
		r.Post("/darajacallback", handlers.Make(handlers.DarajaCallbackHandler))
		r.Post("/fileupload", handlers.Make(handlers.PostUploadFile))

		r.Route("/orders", func(r chi.Router) {
			r.Use(authMiddleware.AuthRedirect)
			r.Get("/available", handlers.Make(handlers.GetAvailableOrders))
			r.Get("/active", handlers.Make(handlers.GetActiveOrders))
			r.Get("/completed", handlers.Make(handlers.GetCompletedOrders))

			r.Post("/{orderID}/take", handlers.Make(handlers.TakeOrder))
			r.Post("/{orderID}/download", handlers.Make(handlers.DownloadOrder))
			r.Post("/{orderID}/cancel", handlers.Make(handlers.CancelOrder))
			r.Post("/{orderID}/complete", handlers.Make(handlers.CompleteOrder))
		})

	})

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r)
}

func public() http.Handler {
	return http.StripPrefix("/public", http.FileServer(http.Dir("public")))
}
