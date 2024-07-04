package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, dashboard.Index())
}

func GetDashBoardContent(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, dashboard.Content())
}
