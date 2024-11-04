package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/internal/config"
	"webapp/internal/models/responses"
	"webapp/internal/pkg"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
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

	response, err := http.Post(fmt.Sprintf("%s/login", config.GetConfig().ApiUrl), "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	token, _ := io.ReadAll(response.Body)

	responses.JSON(w, response.StatusCode, string(token))
}
