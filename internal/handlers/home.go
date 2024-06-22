package handlers

import (
	"net/http"

	home "github.com/mwaurathealex/mbumwa3d/internal/views"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
