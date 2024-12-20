package handlers

import (
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"github.com/mwaurathealex/mbumwa3d/internal/views/home"
)

type HomeHandler struct {
	SessionName string
}

func NewHomeHandler(sessionName string) *HomeHandler {
	return &HomeHandler{
		SessionName: sessionName,
	}
}

func (h *HomeHandler) HandleHome(w http.ResponseWriter, r *http.Request) error {
	if IsHtmx(r) {
		return Render(w, r, home.HomeContent())
	}

	return Render(w, r, home.Index(&store.HomePageParams{}))
}

func (h *HomeHandler) GetUserMenu(w http.ResponseWriter, r *http.Request) error {
	session, _ := gothic.Store.Get(r, h.SessionName)
	_, ok := session.Values["user"].(goth.User)

	if !ok {
		return Render(w, r, components.LoggedOutMenuContent())
	}
	return Render(w, r, components.LoggedInMenuContent())
}

// testing page
func (h *HomeHandler) HandleHello(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Hello())
}
