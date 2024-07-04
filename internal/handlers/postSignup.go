package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/auth"
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
	hasPrinter, err := strconv.ParseBool(r.FormValue("printer-radio"))
	if err != nil {
		hasPrinter = false
	} else {
		hasPrinter = bool(hasPrinter)
	}

	fmt.Println(hasPrinter)

	err = h.userStore.CreateUser(email, password, hasPrinter)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return Render(w, r, auth.SignupError())
	}

	return Render(w, r, auth.SignupSuccess())
}
