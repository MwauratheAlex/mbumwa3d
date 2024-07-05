package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
	"net/http"
	"strconv"
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

func TakeOrder(w http.ResponseWriter, r *http.Request) error {
	// get user id
	user, _ := r.Context().Value(middleware.UserKey).(*store.User)

	// get order id
	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		fmt.Println("Error parsing orderID", err)
		return err
	}

	// get order by id
	orderStore := dbstore.NewOrderStore()
	order, err := orderStore.GetByID(uint(orderID))
	if err != nil {
		fmt.Println("Error Fetching user", err)
		return err
	}
	// update order status
	order.PrintStatus = fmt.Sprint(store.Selected)

	// add user as printer to order
	order.PrinterID = &user.ID

	// save order
	orderStore.Save(order)
	//order-taken-success
	w.Header().Add("HX-Trigger", "orderTakenSuccess")
	return nil
}
