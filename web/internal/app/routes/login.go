package routes

import (
	"net/http"
	"webapp/internal/app/controllers"
)

var loginRoutes = []Route{
	{
		Uri:     "/login",
		Method:  http.MethodGet,
		Handler: controllers.LoadLoginPage,
		Auth:    false,
	},
	{
		Uri:     "/login",
		Method:  http.MethodPost,
		Handler: controllers.LoginPost,
		Auth:    false,
	},
	{
		Uri:     "/logout",
		Method:  http.MethodGet,
		Handler: controllers.Logout,
		Auth:    true,
	},
}
