package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
)

type DashboardHandler struct {
	OrderStore  *dbstore.OrderStore
	SessionName string
}

type DashboardHandlerParams struct {
	OrderStore  *dbstore.OrderStore
	SessionName string
}

func NewDashboardHandler(params DashboardHandlerParams) *DashboardHandler {
	return &DashboardHandler{
		OrderStore:  params.OrderStore,
		SessionName: params.SessionName,
	}
}

func (h *DashboardHandler) HandleDashboard(
	w http.ResponseWriter, r *http.Request) error {
	orders, err := h.OrderStore.GetAvailable()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}
	if IsHtmx(r) {
		return Render(w, r, dashboard.Content(orders))
	}
	return Render(w, r, dashboard.Index(orders))
}

func (h *DashboardHandler) GetAvailable(
	w http.ResponseWriter, r *http.Request) error {
	orders, err := h.OrderStore.GetAvailable()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	return Render(w, r, dashboard.AvailableOrdersTable(orders))
}

func (h *DashboardHandler) GetPrinting(
	w http.ResponseWriter, r *http.Request) error {

	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	orders, err := h.OrderStore.GetPrintOrders(
		uint(userID), store.Printing.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	return Render(w, r, dashboard.ActiveOrdersTable(orders))
}

func (h *DashboardHandler) GetShipping(
	w http.ResponseWriter, r *http.Request) error {

	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	orders, err := h.OrderStore.GetPrintOrders(
		uint(userID), store.Shipping.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}

	return Render(w, r, dashboard.ShippingOrdersTable(orders))
}

func (h *DashboardHandler) GetCompleted(
	w http.ResponseWriter, r *http.Request) error {

	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	orders, err := h.OrderStore.GetPrintOrders(
		uint(userID), store.Completed.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return err
	}
	return Render(w, r, dashboard.CompletedOrdersTable(orders))
}

func (h *DashboardHandler) TakeOrder(w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	order, err := h.OrderStore.GetByID(uint(orderID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	printerID := uint(userID)

	order.PrinterID = &printerID
	order.Status = store.Printing.String()
	err = h.OrderStore.Save(order)

	return err
}

func (h *DashboardHandler) MarkShipping(
	w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order, err := h.OrderStore.GetByID(uint(orderID))
	if err != nil || order.PrinterID == nil || *order.PrinterID != uint(userID) {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order.Status = store.Shipping.String()
	h.OrderStore.Save(order)

	return nil
}

func (h *DashboardHandler) CompleteOrder(
	w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order, err := h.OrderStore.GetByID(uint(orderID))
	if err != nil || order.PrinterID == nil || *order.PrinterID != uint(userID) {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order.Status = store.Completed.String()
	h.OrderStore.Save(order)

	return nil
}

func (h *DashboardHandler) ShipOrder(
	w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order, err := h.OrderStore.GetByID(uint(orderID))
	if err != nil || order.PrinterID == nil || *order.PrinterID != uint(userID) {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order.Status = store.Shipping.String()
	h.OrderStore.Save(order)

	return nil
}

func (h *DashboardHandler) CancelTakenOrder(
	w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *DashboardHandler) DownloadOrder(
	w http.ResponseWriter, r *http.Request) error {
	return nil
}
