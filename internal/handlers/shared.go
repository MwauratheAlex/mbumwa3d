package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Make(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

func IsHtmx(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func GetSessionUser(r *http.Request, sessionName string) (goth.User, error) {
	session, err := gothic.Store.Get(r, sessionName)
	if err != nil {
		return goth.User{}, err
	}

	user, ok := session.Values["user"].(goth.User)

	if !ok {
		return goth.User{}, fmt.Errorf("User is not authenticated! %v", user)
	}

	return user, nil
}

func GetSessionPrintConfig(
	r *http.Request, sessionName string) (store.PrintConfig, error) {

	session, err := gothic.Store.Get(r, sessionName)
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

func ValidatePrintConfig(
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
	if printConfig.Quantity == 0 {
		return fmt.Errorf("Quantity file is required")
	}

	return nil
}
