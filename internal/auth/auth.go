package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
)

type Auth struct {
	AuthConfig  *oauth2.Config
	CookieStore *sessions.CookieStore
}

const oneDay = 86400

const (
	key    = "randomstring" // session secret
	MaxAge = oneDay * 30
	IsProd = false
)

var AuthData *Auth

func NewAuth() {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	url := fmt.Sprintf("%s/auth/google/callback", os.Getenv("BASE_URL"))

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd
	if !IsProd {
		store.Options.SameSite = http.SameSiteLaxMode
	}
	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, url),
	)
}
