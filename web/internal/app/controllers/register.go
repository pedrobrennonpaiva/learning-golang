package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/internal/pkg"
)

func Register(w http.ResponseWriter, r *http.Request) {
	pkg.ExecuteTemplate(w, "register.html", nil)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nickname": r.FormValue("nickname"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:5500/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)
	w.Write([]byte("register post"))
}
