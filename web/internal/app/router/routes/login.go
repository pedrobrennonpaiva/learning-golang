package routes

import (
	"net/http"
	"webapp/internal/app/controllers"
)

var loginRoutes = []Route{
	{
		Uri:     "/",
		Method:  http.MethodGet,
		Handler: controllers.LoadLoginPage,
		Auth:    false,
	},
	{
		Uri:     "/login",
		Method:  http.MethodGet,
		Handler: controllers.LoadLoginPage,
		Auth:    false,
	},
}
