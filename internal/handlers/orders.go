package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
	"github.com/mwaurathealex/mbumwa3d/internal/views/processing"
)

type ToastEventMsg struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

type GetToastPayloadParams struct {
	EventName   string
	Message     string
	Description string
}

func GetToastPayload(params *GetToastPayloadParams) string {
	payload := map[string]ToastEventMsg{
		params.EventName: {
			Message:     params.Message,
			Description: params.Description,
		},
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return string(payloadJson)
}

////////////////////////////////

type OrderHandler struct {
	SessionName string
	OrderStore  *dbstore.OrderStore
	UserStore   *dbstore.UserStore
}

type OrderHandlerParams struct {
	OrderStore  *dbstore.OrderStore
	UserStore   *dbstore.UserStore
	SessionName string
}

func NewOrderHandler(params OrderHandlerParams) *OrderHandler {
	return &OrderHandler{
		SessionName: params.SessionName,
		UserStore:   params.UserStore,
		OrderStore:  params.OrderStore,
	}
}

func (h *OrderHandler) GetProcessing(w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	orders := h.OrderStore.GetNotCompleted(uint(userID))
	if IsHtmx(r) {
		return Render(w, r, processing.Content(orders))
	}
	return Render(w, r, processing.Index(orders))
}

func (h *OrderHandler) GetComplete(w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}

	orders := h.OrderStore.GetCompleted(uint(userID))
	fmt.Println(orders)
	if IsHtmx(r) {
		return Render(w, r, dashboard.AvailableOrdersContent(orders))
	}
	return Render(w, r, dashboard.AvailableOrders(orders))
}

func (h *OrderHandler) MakePayment(w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	order, err := h.UserStore.GetOrder(uint(orderID), uint(userID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	fmt.Println(order)

	return Render(w, r,
		components.SummaryModalContent(
			store.SummaryModalParams{
				PrintContif:    order.PrintConfig,
				IsLoggedInUser: true,
			}),
	)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) error {
	user, err := GetSessionUser(r, h.SessionName)
	userID, err := strconv.ParseInt(user.UserID, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	err = h.OrderStore.Delete(uint(userID), uint(orderID))

	w.Header().Add("HX-Trigger", GetToastPayload(&GetToastPayloadParams{
		EventName:   "OrderEvent",
		Message:     "success",
		Description: "Order deleted",
	}))

	return err
}
