package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
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
	session.Values["user"] = user
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("ERROR adding user to session")
	}

	fmt.Println(userInDb)

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

func (h *AuthHandler) GetSessionUser(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, h.SessionName)
	if err != nil {
		return goth.User{}, err
	}

	user := session.Values["user"]
	if user == nil {
		return goth.User{}, fmt.Errorf("User is not authenticated! %v", user)
	}

	return user.(goth.User), nil
}

func (h *AuthHandler) RequreAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := h.GetSessionUser(r)
		if err != nil {
			log.Println("User is not authenticated!")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		log.Printf("user is authenticated! user: %v!", session.Name)
		handlerFunc(w, r)

	}
}
