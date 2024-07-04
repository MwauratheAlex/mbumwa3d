package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/finished"
)

func HandleFinished(w http.ResponseWriter, r *http.Request) error {
	if IsHtmx(r) {
		// return content
	}
	return Render(w, r, finished.Index())
}
