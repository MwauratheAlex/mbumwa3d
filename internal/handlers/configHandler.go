package handlers

import (
	"net/http"
	"strconv"

	"github.com/markbates/goth/gothic"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"github.com/mwaurathealex/mbumwa3d/internal/views/printSummary"
)

type PrintSummaryHandler struct {
	SessionName string
	UserStore   store.UserStore
}

type PrintSummaryHandlerParams struct {
	SessionName string
	UserStore   store.UserStore
}

func NewPrintSummaryHandler(params PrintSummaryHandlerParams) *PrintSummaryHandler {
	if params.SessionName == "" {
		panic("Session Name is requred in PrintConfigHandler")
	}
	return &PrintSummaryHandler{
		SessionName: params.SessionName,
		UserStore:   params.UserStore,
	}
}

func (h *PrintSummaryHandler) HandlePrintSummary(w http.ResponseWriter, r *http.Request) error {

	printConfig, _ := GetSessionPrintConfig(r, h.SessionName)
	qty, _ := strconv.ParseInt(r.FormValue("quantity"), 10, 64)

	if r.Method == "POST" {
		printConfig.Technology = r.FormValue("technology")
		printConfig.Material = r.FormValue("material")
		printConfig.Color = r.FormValue("color")
		printConfig.Quantity = int(qty)
	}

	err := ValidatePrintConfig(&printConfig)

	errorEventPayload := &GetToastPayloadParams{
		EventName: "FileConfigUploadEvent",
		Message:   "error",
	}

	if err != nil {
		errorEventPayload.Description = err.Error()
		w.Header().Add("HX-Trigger", GetToastPayload(errorEventPayload))
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	session, _ := gothic.Store.Get(r, h.SessionName)
	session.Values["print_config"] = printConfig
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic("Error saving printConfig to Session")
	}

	_, err = GetSessionUser(r, h.SessionName)
	isLoggedIn := err == nil

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	if r.Method == "GET" {
		if IsHtmx(r) {
			return Render(w, r,
				components.SummaryModalContent(
					store.SummaryModalParams{
						PrintContif:    printConfig,
						IsLoggedInUser: isLoggedIn,
					}),
			)
		}

		return Render(w, r,
			printSummary.PrintSummaryPage(
				store.SummaryModalParams{
					PrintContif:    printConfig,
					IsLoggedInUser: isLoggedIn,
				}),
		)
	}

	return Render(w, r, components.SummaryModalContent(store.SummaryModalParams{
		IsLoggedInUser: isLoggedIn,
		PrintContif:    printConfig,
	}))
}
