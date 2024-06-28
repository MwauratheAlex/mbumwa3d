package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
	"github.com/mwaurathealex/mbumwa3d/internal/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}

func GetUserMenu(w http.ResponseWriter, r *http.Request) error {
	_, ok := r.Context().Value(middleware.UserKey).(*store.User)
	fmt.Println(ok)
	if ok {
		return Render(w, r, components.LoggedInUserMenu())
	}

	return Render(w, r, components.LoggedOutUserMenu())
}
