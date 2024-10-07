package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
)

type DashboardHandler struct {
	OrderStore *dbstore.OrderStore
}

type DashboardHandlerParams struct {
	OrderStore *dbstore.OrderStore
}

func NewDashboardHandler(params DashboardHandlerParams) *DashboardHandler {
	return &DashboardHandler{
		OrderStore: params.OrderStore,
	}
}

func (h *DashboardHandler) HandleDashboard(
	w http.ResponseWriter, r *http.Request) error {
	orders, err := h.OrderStore.GetAvailable()
	fmt.Println(err)
	if IsHtmx(r) {
		return Render(w, r, dashboard.Content(orders))
	}
	return Render(w, r, dashboard.Index(orders))
}

func (h *DashboardHandler) GetAvailable(
	w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *DashboardHandler) GetActive(
	w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *DashboardHandler) GetCompleted(
	w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *DashboardHandler) TakeOrder(w http.ResponseWriter, r *http.Request) error {
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

func (h *DashboardHandler) CompleteOrder(
	w http.ResponseWriter, r *http.Request) error {
	return nil
}
