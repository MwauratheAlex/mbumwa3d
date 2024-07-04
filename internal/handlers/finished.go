package handlers

import (
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/finished"
)

func HandleFinished(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if ok == false {
		w.Header().Add("HX-Redirect", "/login")
		return nil
	}

	orderStore := dbstore.NewOrderStore()
	completedOrders := orderStore.GetCompleted(user.ID)

	if IsHtmx(r) {
		return Render(w, r, finished.Content(completedOrders))
	}
	return Render(w, r, finished.Index(completedOrders))
}
