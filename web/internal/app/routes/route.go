package routes

import (
	"net/http"
	"webapp/internal/app/middlewares"

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
	routes = append(routes, usersRoutes...)
	routes = append(routes, homeRoutes...)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.Auth {
			router.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Handler))).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
