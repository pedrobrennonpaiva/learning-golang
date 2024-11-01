package app

import (
	"webapp/internal/app/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	router := mux.NewRouter()

	return routes.Configure(router)
}
