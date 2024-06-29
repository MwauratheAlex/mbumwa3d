package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/views/auth"
	"golang.org/x/crypto/bcrypt"
)

type PostLoginHandler struct {
	userStore  store.UserStore
	CookieName string
}

type Claims struct {
	Sub uint
	jwt.RegisteredClaims
}

type PostLoginHandlerParams struct {
	UserStore store.UserStore
}

func NewPostLoginHandler(params PostLoginHandlerParams) http.HandlerFunc {
	return Make((&PostLoginHandler{
		userStore:  params.UserStore,
		CookieName: "Authorization",
	}).PostLogin)

}

func (h *PostLoginHandler) PostLogin(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.userStore.GetUser(email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return Render(w, r, auth.LoginError())
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash), []byte(password),
	); err != nil {
		fmt.Println(err, password, user.Email)
		w.WriteHeader(http.StatusUnauthorized)
		return Render(w, r, auth.LoginError())
	}

	// jwt stuff
	expiration := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Sub: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// cookie stuff - will be the generated jwt token
	authSecret := os.Getenv("AUTH_SECRET")
	cookieValue, err := token.SignedString([]byte(authSecret))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return Render(w, r, auth.LoginError())
	}

	cookie := http.Cookie{
		Name:     h.CookieName,
		Value:    cookieValue,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	// successful login
	w.Header().Add("HX-Redirect", "/")
	w.Header().Add("HX-Trigger", "login_success")
	w.WriteHeader(http.StatusOK)
	return nil
}
