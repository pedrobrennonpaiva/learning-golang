package routes

import (
	"net/http"
	"webapp/internal/app/controllers"
)

var usersRoutes = []Route{
	{
		Uri:     "/register",
		Method:  http.MethodGet,
		Handler: controllers.Register,
		Auth:    false,
	},
	{
		Uri:     "/register",
		Method:  http.MethodPost,
		Handler: controllers.RegisterPost,
		Auth:    false,
	},
}
