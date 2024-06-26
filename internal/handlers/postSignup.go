package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/components"
)

type PostSignupHandler struct {
	userStore store.UserStore
}

type PostSignupHandlerParams struct {
	UserStore store.UserStore
}

func NewPostSignupHandler(params PostSignupHandlerParams) http.HandlerFunc {
	return Make((&PostSignupHandler{
		userStore: params.UserStore,
	}).PostSignup)
}

func (h *PostSignupHandler) PostSignup(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println("Email ", email, " Password: ", password)

	err := h.userStore.CreateUser(email, password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	return Render(w, r, components.SuccessMsg("Signup suceessful"))
}
