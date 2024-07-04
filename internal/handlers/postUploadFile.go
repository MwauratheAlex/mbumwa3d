package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
)

func PostUploadFile(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, components.UploadForm())
}
