package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/auth"
	"github.com/mwaurathealex/mbumwa3d/internal/handlers"
	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"github.com/mwaurathealex/mbumwa3d/internal/payment"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

var Environment string = "dev"

func init() {
	// register stuff to cookie store
	gob.Register(store.PrintConfig{})

	switch Environment {
	case "docker-prod":
		os.Setenv("env", "production")
		break
	case "dev":
		initializers.LoadEnvVariables()
	}
	initializers.ConnectToDB()
}

func main() {

	r := chi.NewMux()

	auth.NewAuth()
	port := os.Getenv("PORT")

	userStore := dbstore.NewUserStore()
	fileStore := dbstore.NewFileStore()
	orderStore := dbstore.NewOrderStore()

	sessionName := "user-session"
	homeHandler := handlers.NewHomeHandler(sessionName)

	authHandler := handlers.NewAuthHandler(handlers.AuthHandlerParams{
		UserStore:   userStore,
		SessionName: sessionName},
	)

	fileHandler := handlers.NewFileHandler(handlers.FileHandlerParams{
		SessionName: sessionName,
		FileStore:   fileStore,
	})
	printSummaryHandler := handlers.NewPrintSummaryHandler(
		handlers.PrintSummaryHandlerParams{
			SessionName: sessionName,
			UserStore:   userStore,
		},
	)
	paymentProcessor := payment.NewPaymentProcessor()
	paymentHandler := handlers.NewPaymentHandler(handlers.PaymentHandlerParams{
		PaymentProcessor: paymentProcessor,
		SessionName:      sessionName,
		OrderStore:       orderStore,
	})

	orderHandler := handlers.NewOrderHandler(handlers.OrderHandlerParams{
		SessionName: sessionName,
		UserStore:   userStore,
		OrderStore:  orderStore,
	})

	r.Group(func(r chi.Router) {
		//r.Use(
		//middleware.TextHTMLMiddleware,
		// authMiddleware.AddUserToContext,
		//)
		r.Handle("/*", public())
		r.Get("/", handlers.Make(homeHandler.HandleHome))

		r.Get("/auth/{provider}/callback", handlers.Make(authHandler.AuthCallback))
		r.Get("/auth/{provider}", handlers.Make(authHandler.BeginAuth))

		r.Post("/file", fileHandler.Post)
		r.Post("/print-summary", handlers.Make(
			printSummaryHandler.HandlePrintSummary,
		))
		r.Get("/print-summary", handlers.Make(
			printSummaryHandler.HandlePrintSummary,
		))
		r.Post("/payment", handlers.Make(paymentHandler.Post))
		r.Post("/darajacallback", handlers.Make(paymentHandler.DarajaCallback))
		r.Post("/payment-status-callback", handlers.Make(
			paymentHandler.DarajaPaymentStatusCallback,
		))
		r.Post("/payment-confirmation", paymentHandler.PaymentNotificationCallback)

		r.Route("/orders", func(r chi.Router) {
			r.Get("/processing", handlers.Make(orderHandler.GetProcessing))
			r.Get("/complete", handlers.Make(orderHandler.GetComplete))

			// for user
			r.Post("/{orderID}/make-payment", handlers.Make(
				orderHandler.MakePayment))
			r.Delete("/{orderID}/delete", handlers.Make(orderHandler.DeleteOrder))

			// for printer
			r.Post("/{orderID}/take", func(w http.ResponseWriter, r *http.Request) {})
			r.Post("/{orderID}/download", func(w http.ResponseWriter, r *http.Request) {})
			r.Post("/{orderID}/printer-cancel", func(w http.ResponseWriter, r *http.Request) {})
			r.Post("/{orderID}/complete", func(w http.ResponseWriter, r *http.Request) {})
		})

		///////////

		r.Get("/processing", handlers.Make(handlers.HandleProcessing))
		r.Get("/usermenu", handlers.Make(homeHandler.GetUserMenu))
		r.Get("/dashboard", handlers.Make(handlers.HandleDashboard))

		// r.Route("/orders", func(r chi.Router) {
		// 	// r.Use(authMiddleware.AuthRedirect)
		// 	r.Get("/available", handlers.Make(handlers.GetAvailableOrders))
		// 	r.Get("/active", handlers.Make(handlers.GetActiveOrders))
		// 	r.Get("/completed", handlers.Make(handlers.GetCompletedOrders))

		// 	r.Post("/{orderID}/take", handlers.Make(handlers.TakeOrder))
		// 	r.Post("/{orderID}/download", handlers.Make(handlers.DownloadOrder))
		// 	r.Post("/{orderID}/cancel", handlers.Make(handlers.CancelOrder))
		// 	r.Post("/{orderID}/complete", handlers.Make(handlers.CompleteOrder))
		// })

	})

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r)
}

func public() http.Handler {
	return http.StripPrefix("/public", http.FileServer(http.Dir("public")))
}
