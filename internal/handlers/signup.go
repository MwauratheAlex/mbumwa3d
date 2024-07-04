package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/auth"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) error {
	if IsHtmx(r) {
		return Render(w, r, auth.SignupContent())
	}
	return Render(w, r, auth.Signup())
}
