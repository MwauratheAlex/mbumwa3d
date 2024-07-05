package handlers

import (
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
	"net/http"
)

func GetAvailableOrders(w http.ResponseWriter, r *http.Request) error {
	orderStore := dbstore.NewOrderStore()
	availableOrders := orderStore.GetPrintAvailable()
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
