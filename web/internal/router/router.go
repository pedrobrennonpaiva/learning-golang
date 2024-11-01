package router

import "github.com/gorilla/mux"

func Generate() *mux.Router {
	router := mux.NewRouter()
	return router
}
