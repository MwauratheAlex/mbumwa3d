package handlers

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth"
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

	printConfig, _ := h.GetSessionPrintConfig(r)

	if r.Method == "POST" {
		printConfig.Technology = r.FormValue("technology")
		printConfig.Material = r.FormValue("material")
		printConfig.Color = r.FormValue("color")
		printConfig.Quantity = r.FormValue("quantity")
	}

	err := h.ValidateFormData(&printConfig)

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

	_, err = h.GetSessionUser(r)
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

func (h *PrintSummaryHandler) ValidateFormData(
	printConfig *store.PrintConfig) error {

	if printConfig.FileID == "" || printConfig.FileVolume == 0 {
		return fmt.Errorf("STL file is required")
	}
	if printConfig.Technology == "" {
		return fmt.Errorf("Technology file is required")
	}
	if printConfig.Material == "" {
		return fmt.Errorf("Material file is required")
	}
	if printConfig.Color == "" {
		return fmt.Errorf("Color file is required")
	}
	if printConfig.Quantity == "" {
		return fmt.Errorf("Quantity file is required")
	}

	return nil
}

func (h *PrintSummaryHandler) GetSessionUser(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, h.SessionName)
	if err != nil {
		return goth.User{}, err
	}

	user, ok := session.Values["user"].(goth.User)

	if !ok {
		return goth.User{}, fmt.Errorf("User is not authenticated! %v", user)
	}

	return user, nil
}

func (h *PrintSummaryHandler) GetSessionPrintConfig(
	r *http.Request) (store.PrintConfig, error) {

	session, err := gothic.Store.Get(r, h.SessionName)
	if err != nil {
		return store.PrintConfig{}, err
	}

	printConfig, ok := session.Values["print_config"].(store.PrintConfig)

	if !ok {
		return store.PrintConfig{}, fmt.Errorf(
			"Print Config not set! %v", printConfig,
		)
	}

	return printConfig, nil
}
