package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/internal/models/responses"
	"webapp/internal/pkg"
)

func Register(w http.ResponseWriter, r *http.Request) {
	pkg.ExecuteTemplate(w, "register.html", nil)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nickname": r.FormValue("nickname"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5500/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
