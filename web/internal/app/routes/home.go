package routes

import (
	"net/http"
	"webapp/internal/app/controllers"
)

var homeRoutes = []Route{
	{
		Uri:     "/",
		Method:  http.MethodGet,
		Handler: controllers.Home,
		Auth:    true,
	},
	{
		Uri:     "/home",
		Method:  http.MethodGet,
		Handler: controllers.Home,
		Auth:    true,
	},
}
