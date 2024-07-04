package handlers

import (
	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/processing"
	"net/http"
)

func HandleProcessing(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if ok == false {
		w.Header().Add("HX-Redirect", "/login")
		return nil
	}

	orderStore := dbstore.NewOrderStore()
	processingOrders := orderStore.GetNotCompleted(user.ID)

	if IsHtmx(r) {
		return Render(w, r, processing.Content(processingOrders))
	}
	return Render(w, r, processing.Index(processingOrders))
}
