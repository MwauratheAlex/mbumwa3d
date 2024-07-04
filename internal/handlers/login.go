package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/views/auth"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	if IsHtmx(r) {
		return Render(w, r, auth.LoginContent())
	}
	return Render(w, r, auth.Login())
}
