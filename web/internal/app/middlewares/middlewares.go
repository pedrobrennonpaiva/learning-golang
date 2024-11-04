package middlewares

import (
	"log"
	"net/http"
	"time"
	"webapp/internal/pkg/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authResponse, err := cookies.Read(r)
		if err != nil ||
			authResponse.ID == "" ||
			authResponse.Token == "" ||
			authResponse.ExpiresAt == 0 ||
			time.Unix(authResponse.ExpiresAt, 0).Before(time.Now()) {
			http.Redirect(w, r, "/login", 302)
		} else {
			next(w, r)
		}
	}
}
