package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webapp/internal/config"
	"webapp/internal/pkg"
	"webapp/internal/pkg/cookies"
	"webapp/internal/pkg/responses"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie.Token != "" && cookie.ExpiresAt > time.Now().Unix() {
		http.Redirect(w, r, "/", 302)
		return
	}

	pkg.ExecuteTemplate(w, "login.html", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.GetConfig().ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	var authResponse cookies.AuthResponse
	if err = json.NewDecoder(response.Body).Decode(&authResponse); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	if err = cookies.Save(w, authResponse); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	responses.JSON(w, response.StatusCode, authResponse)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
