package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
)

type AuthHandler struct {
	UserStore   store.UserStore
	SessionName string
}

type AuthHandlerParams struct {
	UserStore   store.UserStore
	SessionName string
}

func NewAuthHandler(params AuthHandlerParams) *AuthHandler {
	return &AuthHandler{
		UserStore:   params.UserStore,
		SessionName: params.SessionName,
	}
}

func (h *AuthHandler) BeginAuth(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
	return nil
}

func (h *AuthHandler) AuthCallback(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	userInDb, err := h.UserStore.GetOrCreate(&store.User{
		Email:    user.Email,
		Name:     user.Name,
		PhotoUrl: user.AvatarURL,
	})

	session, _ := gothic.Store.Get(r, h.SessionName)
	session.Values["user_id"] = userInDb.ID
	session.Save(r, w)

	//	 TODO:
	//
	// 3. fileId  = cookie[fileId]
	// 5. if !file, then return; show toast, login success
	//
	// 6. else we show the summary of config?????? from local storage, so we can just send down an event (form-after-login???)
	// 7. Edit or continue to payment.

	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	return nil
}
