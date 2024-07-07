package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
)

var NonceKey = "nonces"

type Nonces struct {
	Htmx            string
	ResponseTargets string
	Tw              string
	HtmxCSSHash     string
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonceSet := Nonces{
			Htmx:            generateRandomString(16),
			ResponseTargets: generateRandomString(16),
			Tw:              generateRandomString(16),
			HtmxCSSHash:     "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg=",
		}
		// set nonces in context
		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)
		// insert nonces into content security policy header
		cspHeader := fmt.Sprintf(
			"default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' '%s';",
			nonceSet.Htmx,
			nonceSet.ResponseTargets,
			nonceSet.Tw,
			nonceSet.HtmxCSSHash,
		)
		w.Header().Set("Content-Security-Policy", cspHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// get the Nonce from the context, it is a struct called Nonces,
// so we can get the nonce we need by the key, i.e. HtmxNonce
func GetNonces(ctx context.Context) Nonces {
	nonceSet := ctx.Value(NonceKey)
	if nonceSet == nil {
		log.Fatal("error getting nonce set - is nil")
	}

	nonces, ok := nonceSet.(Nonces)

	if !ok {
		log.Fatal("error getting nonce set - not ok")
	}

	return nonces
}

func GetHtmxNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)

	return nonceSet.Htmx
}

func GetResponseTargetsNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.ResponseTargets
}

func GetTwNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Tw
}

func TextHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

type AuthMiddleware struct {
	CookieName string
	Userstore  store.UserStore
}

func NewAuthMiddleware(cookieName string, userstore store.UserStore) *AuthMiddleware {
	return &AuthMiddleware{
		CookieName: cookieName,
		Userstore:  userstore,
	}
}

var UserKey string = "user"

func (m *AuthMiddleware) AddUserToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(m.CookieName)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		cookieValue := cookie.Value

		// validate cookie
		authSecret := os.Getenv("AUTH_SECRET")
		token, err := jwt.Parse(cookieValue, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(authSecret), nil
		})
		if err != nil {
			fmt.Println("error getting the token from cookie")
			next.ServeHTTP(w, r)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				next.ServeHTTP(w, r)
				return
			}
			userId := claims["Sub"].(float64)
			user, err := m.Userstore.GetUserById(uint(userId))

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Println("claims not okay")
			next.ServeHTTP(w, r)
		}

	})
}

func (m *AuthMiddleware) AuthRedirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(UserKey).(*store.User)

		if ok == false {
			w.Header().Add("HX-Redirect", "/login")
			return
		}
		next.ServeHTTP(w, r)

	})

}

func GetUser(ctx context.Context) *store.User {
	user := ctx.Value(UserKey)
	if user == nil {
		return nil
	}

	return user.(*store.User)
}
