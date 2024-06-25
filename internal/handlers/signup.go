package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/auth"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Signup())
}
