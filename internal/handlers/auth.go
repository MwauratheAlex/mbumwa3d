package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func GetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(w, err)
		return err
	}
	fmt.Println(user)
	return nil
}

func OauthLoginHandler(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")

	fmt.Println("Hello world")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)

	return nil
}
