package pkg

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("internal/web/*html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
