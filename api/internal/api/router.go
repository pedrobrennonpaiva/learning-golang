package api

import (
	"golang-api/internal/api/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	router := mux.NewRouter()

	return routes.Configure(router)
}
