package routes

import (
	"golang-api/internal/api/middlewares"
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
	routes := routesUser
	routes = append(routes, routesLogin...)
	routes = append(routes, routesPost...)

	for _, route := range routes {

		if route.Auth {
			router.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	return router
}
