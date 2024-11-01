package controllers

import (
	"net/http"
	"webapp/internal/pkg"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	pkg.ExecuteTemplate(w, "login.html", nil)
}
