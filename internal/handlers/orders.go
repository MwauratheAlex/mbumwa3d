package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
)

func GetAvailableOrders(w http.ResponseWriter, r *http.Request) error {
	orderStore := dbstore.NewOrderStore()
	availableOrders := orderStore.GetPrintAvailable()
	for _, order := range availableOrders {
		fmt.Print(order.File.FileName)
	}
	if IsHtmx(r) {
		return Render(w, r, dashboard.AvailableOrdersContent(availableOrders))
	}
	return Render(w, r, dashboard.AvailableOrders(availableOrders))
}
func GetActiveOrders(w http.ResponseWriter, r *http.Request) error {

	user, _ := r.Context().Value(middleware.UserKey).(*store.User)

	orderStore := dbstore.NewOrderStore()
	activeOrders := orderStore.GetPrintActive(user.ID)

	if IsHtmx(r) {
		return Render(w, r, dashboard.ActiveOrdersContent(activeOrders))
	}
	return Render(w, r, dashboard.ActiveOrders(activeOrders))
}

func GetCompletedOrders(w http.ResponseWriter, r *http.Request) error {
	user, _ := r.Context().Value(middleware.UserKey).(*store.User)

	orderStore := dbstore.NewOrderStore()
	activeOrders := orderStore.GetPrintCompleted(user.ID)

	if IsHtmx(r) {
		return Render(w, r, dashboard.CompletedOrdersContent(activeOrders))
	}
	return Render(w, r, dashboard.CompletedOrders(activeOrders))
}
