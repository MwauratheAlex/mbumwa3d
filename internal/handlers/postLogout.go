package handlers

import (
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"net/http"
	"time"
)

func PostLogout(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
	return Render(w, r, components.LoggedOutUserMenu())
}
