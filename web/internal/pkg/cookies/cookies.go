package cookies

import (
	"net/http"
	"time"
	"webapp/internal/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New([]byte(config.GetConfig().HashKey), []byte(config.GetConfig().BlockKey))
}

func Save(w http.ResponseWriter, authResponse AuthResponse) error {
	encoded, err := s.Encode("authData", authResponse)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "authData",
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (AuthResponse, error) {
	cookie, err := r.Cookie("authData")
	if err != nil {
		return AuthResponse{}, err
	}

	var auth AuthResponse
	if err = s.Decode("authData", cookie.Value, &auth); err != nil {
		return AuthResponse{}, err
	}

	return auth, nil
}

func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "authData",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
