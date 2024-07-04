package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/dashboard"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) error {
	if IsHtmx(r) {
		return Render(w, r, dashboard.Content())
	}
	return Render(w, r, dashboard.Index())
}
