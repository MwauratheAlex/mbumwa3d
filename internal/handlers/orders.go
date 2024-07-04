package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
)

func GetActiveOrders(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Active")
	return Render(w, r, dashboard.ActiveOrders())
}
func GetAvailableOrders(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Available")
	return Render(w, r, dashboard.AvailableOrders())
}
func GetCompletedOrders(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Completed")
	return Render(w, r, dashboard.CompletedOrders())
}
