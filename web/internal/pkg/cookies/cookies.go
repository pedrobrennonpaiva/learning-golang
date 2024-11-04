package cookies

import (
	"net/http"
	"webapp/internal/config"
	"webapp/internal/models"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New([]byte(config.GetConfig().HashKey), []byte(config.GetConfig().BlockKey))
}

func Save(w http.ResponseWriter, authResponse models.AuthResponse) error {
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

func Read(r *http.Request) (models.AuthResponse, error) {
	cookie, err := r.Cookie("authData")
	if err != nil {
		return models.AuthResponse{}, err
	}

	var auth models.AuthResponse
	if err = s.Decode("authData", cookie.Value, &auth); err != nil {
		return models.AuthResponse{}, err
	}

	return auth, nil
}
