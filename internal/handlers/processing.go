package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/processing"
)

func HandleProcessing(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, processing.Index())
}

func GetProcessingContent(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, processing.Content())
}
