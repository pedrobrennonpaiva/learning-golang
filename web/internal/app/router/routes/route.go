package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
	Auth    bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
