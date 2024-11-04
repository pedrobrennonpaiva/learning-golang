package cookies

import (
	"net/http"
	"webapp/internal/config"
	"webapp/internal/models/responses"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New([]byte(config.GetConfig().HashKey), []byte(config.GetConfig().BlockKey))
}

func Save(w http.ResponseWriter, authResponse responses.AuthResponse) error {
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
